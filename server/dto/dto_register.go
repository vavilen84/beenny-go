package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
)

type Register struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Register) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", validation.CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func (Register) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioRegister: validation.FieldRules{
			"Nickname": "min=3,max=255,required",
			"Email":    "min=3,max=255,email,required",
			"Password": "min=8,max=255,required,customPasswordValidator",
		},
	}
}
