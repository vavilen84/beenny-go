package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"testing"
)

func Test_PasswordValidation_min(t *testing.T) {
	u := User{
		Password: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	if err != nil {
		log.Fatalln(err)
		return
	}
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "User", "Password", "8"), v["Password"][0].Message)
}

func Test_PasswordValidation_customValidation_1(t *testing.T) {
	u := User{
		Password: "12345678",
	}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	if err != nil {
		log.Fatalln(err)
		return
	}
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "User"), v["Password"][0].Message)
}

func Test_PasswordValidation_customValidation_2(t *testing.T) {
	u := User{
		Password: "12345678l",
	}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	if err != nil {
		log.Fatalln(err)
		return
	}
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "User"), v["Password"][0].Message)
}

func Test_PasswordValidation_customValidation_3(t *testing.T) {
	u := User{
		Password: "12345678lT",
	}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	if err != nil {
		log.Fatalln(err)
		return
	}
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "User"), v["Password"][0].Message)
}

func Test_PasswordValidation_customValidation_4(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	if err != nil {
		log.Fatalln(err)
		return
	}
	assert.Nil(t, err)
}
