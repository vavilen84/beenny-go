package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/interfaces"
)

// should be passed ptr to model m otherwise - func will panic
func ValidateByScenario(scenario Scenario, m interfaces.Model) error {
	validate := m.GetValidator().(*validator.Validate)
	validationMap := m.GetValidationRules().(ScenarioRules)
	var result error
	errs := make(Errors)
	data := helpers.StructToMap(m)
	for fieldName, validation := range validationMap[scenario] {
		field, ok := data[fieldName]
		if !ok {
			helpers.LogFatal(fmt.Sprintf("Field not found: %s", fieldName))
		}
		err := validate.Var(field, string(validation))
		if err != nil {
			if _, ok := errs[fieldName]; !ok {
				errs[fieldName] = make([]FieldError, 0)
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
				errs[fieldName] = append(errs[fieldName], validationError)
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return result
}
