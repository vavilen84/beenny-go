package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/validation"
)

type ForgotPassword struct {
	Email string `json:"email"`
}

func (ForgotPassword) GetValidator() interface{} {
	v := validator.New()
	return v
}

func (ForgotPassword) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioForgotPassword: validation.FieldRules{
			"Email": "required,max=255,email",
		},
	}
}
