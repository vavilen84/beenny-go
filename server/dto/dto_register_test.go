package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/validation"
	"log"
	"testing"
)

func Test_Register_notOk_1(t *testing.T) {
	u := Register{}
	err := validation.ValidateByScenario(constants.ScenarioRegister, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg), v["Email"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "FirstName"), v["FirstName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "LastName"), v["LastName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Photo"), v["Photo"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Gender"), v["Gender"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Timezone"), v["Timezone"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Birthday"), v["Birthday"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "AgreeTerms"), v["AgreeTerms"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "ConfirmPassword"), v["ConfirmPassword"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Password"), v["Password"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "CurrentCountry"), v["CurrentCountry"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "CountryOfBirth"), v["CountryOfBirth"][0].Message)
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
	err := validation.ValidateByScenario(constants.ScenarioRegister, u)
	assert.Nil(t, err)
}
