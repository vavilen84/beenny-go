package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
)

type ResetPassword struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func (ResetPassword) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", validation.CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func (ResetPassword) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioResetPassword: validation.FieldRules{
			"Token":       "max=255,required",
			"NewPassword": "max=255,required,customPasswordValidator",
		},
	}
}
