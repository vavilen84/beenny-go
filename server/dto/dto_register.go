package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
)

type Register struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	CurrentCountry  string `json:"currentCountry"`
	CountryOfBirth  string `json:"countryOfBirth"`
	Gender          string `json:"gender"`
	Timezone        string `json:"timezone"`
	Birthday        string `json:"birthday"`
	AgreeTerms      bool   `json:"agreeTerms"`
	Password        string `json:"password"`
	Photo           string `json:"photo"`
	ConfirmPassword string `json:"confirmPassword"`
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
			"FirstName":       "max=255,required",
			"LastName":        "max=255,required",
			"Email":           "max=255,email,required",
			"CurrentCountry":  "max=2,required",
			"CountryOfBirth":  "max=2,required",
			"Gender":          "max=10,required",
			"Timezone":        "max=255,required",
			"Birthday":        "required",
			"AgreeTerms":      "required",
			"Photo":           "max=255,required",
			"Password":        "max=255,required,customPasswordValidator",
			"ConfirmPassword": "max=255,required,customPasswordValidator",
		},
	}
}
