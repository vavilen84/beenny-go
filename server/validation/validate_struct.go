package validation

import (
	"github.com/go-playground/validator/v10"
)

func Validate(model interface{}) error {
	v := validator.New()
	err := v.Struct(model)
	if err != nil {
		result := make(Errors)
		result.setStructValidationErrors(err, model)
		return result
	}
	return nil
}
