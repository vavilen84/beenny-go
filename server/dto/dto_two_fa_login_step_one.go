package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
)

type TwoFaLoginStepOne struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (TwoFaLoginStepOne) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", validation.CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func (TwoFaLoginStepOne) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioTwoFaLoginStepOne: validation.FieldRules{
			"Email":    "min=3,max=255,email,required",
			"Password": "min=8,max=255,required,customPasswordValidator",
		},
	}
}
