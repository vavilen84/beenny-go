package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/interfaces"
)

// should be passed ptr to model m otherwise - func will panic
func ValidateByScenario(scenario Scenario, m interfaces.Model) (errs Errors) {
	validate := m.GetValidator().(*validator.Validate)
	validationMap := m.GetValidationRules().(ScenarioRules)
	data := helpers.StructToMap(m)
	for fieldName, validation := range validationMap[scenario] {
		field, ok := data[fieldName]
		if !ok {
			helpers.LogFatal(fmt.Sprintf("Field not found: %s", fieldName))
		}
		err := validate.Var(field, string(validation))
		if err != nil {
			if errs == nil {
				errs = make(Errors, 0)
			}
			for _, e := range err.(validator.ValidationErrors) {
				validationError := FieldError{
					Name:  getType(m),
					Tag:   e.Tag(),
					Field: fieldName,
					Value: fmt.Sprintf("%v", e.Value()),
					Param: e.Param(),
				}
				validationError.setErrorMessage()
				errs = append(errs, validationError.Message)
			}
		}
	}
	return errs
}
