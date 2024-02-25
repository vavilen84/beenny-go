package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_Register_notOk_1(t *testing.T) {
	u := Register{}
	errs := validation.ValidateByScenario(constants.ScenarioRegister, u)

	mustHaveErrors := []string{
		fmt.Sprintf(constants.EmailErrorMsg, "Email"),
		fmt.Sprintf(constants.RequiredErrorMsg, "FirstName"),
		fmt.Sprintf(constants.RequiredErrorMsg, "LastName"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Photo"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Gender"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Timezone"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Birthday"),
		fmt.Sprintf(constants.RequiredErrorMsg, "AgreeTerms"),
		fmt.Sprintf(constants.RequiredErrorMsg, "ConfirmPassword"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
		fmt.Sprintf(constants.RequiredErrorMsg, "CurrentCountry"),
		fmt.Sprintf(constants.RequiredErrorMsg, "CountryOfBirth"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_Register_ok(t *testing.T) {
	u := Register{
		FirstName:       "John",
		LastName:        "Dou",
		Email:           "email@example.com",
		CurrentCountry:  "UA",
		CountryOfBirth:  "UA",
		Gender:          constants.GenderMale,
		Timezone:        "US/Arizona",
		Birthday:        "1984-01-23",
		AgreeTerms:      true,
		Password:        "12345678lT*",
		ConfirmPassword: "12345678lT*",
		Photo:           "/2024/01/23/s09d8fs09dfu.jpg",
	}
	errs := validation.ValidateByScenario(constants.ScenarioRegister, u)
	assert.Nil(t, errs)
}
