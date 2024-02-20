package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"testing"
)

func Test_PasswordValidation_min(t *testing.T) {
	u := ChangePassword{
		OldPassword: "",
		NewPassword: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "ChangePassword", "OldPassword", "8"), v["OldPassword"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "ChangePassword", "NewPassword", "8"), v["NewPassword"][0].Message)
}

func Test_PasswordValidation_customValidation_1(t *testing.T) {
	u := ChangePassword{
		OldPassword: "12345678",
		NewPassword: "12345678",
	}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ChangePassword"), v["OldPassword"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ChangePassword"), v["NewPassword"][0].Message)
}

func Test_PasswordValidation_customValidation_2(t *testing.T) {
	u := ChangePassword{
		OldPassword: "12345678l",
		NewPassword: "12345678l",
	}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ChangePassword"), v["OldPassword"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ChangePassword"), v["NewPassword"][0].Message)
}

func Test_PasswordValidation_customValidation_3(t *testing.T) {
	u := ChangePassword{
		OldPassword: "12345678lT",
		NewPassword: "12345678lT",
	}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ChangePassword"), v["OldPassword"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ChangePassword"), v["NewPassword"][0].Message)
}

func Test_PasswordValidation_customValidation_4(t *testing.T) {
	u := ChangePassword{
		OldPassword: "12345678lT*",
		NewPassword: "12345678lT*",
	}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	assert.Nil(t, err)
}
