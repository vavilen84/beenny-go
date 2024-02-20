package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
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
			"Email": "min=3,max=255,email,required",
		},
	}
}
