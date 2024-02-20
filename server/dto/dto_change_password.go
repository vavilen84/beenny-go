package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
)

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (ChangePassword) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", validation.CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func (ChangePassword) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioChangePassword: validation.FieldRules{
			"OldPassword": "min=8,max=255,required,customPasswordValidator",
			"NewPassword": "min=8,max=255,required,customPasswordValidator",
		},
	}
}
