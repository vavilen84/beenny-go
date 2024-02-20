package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"testing"
)

func Test_DTO_resetPassword_notOk_1(t *testing.T) {
	u := ResetPassword{
		Token:       "",
		NewPassword: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "ResetPassword", "Token", "6"), v["Token"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "ResetPassword", "NewPassword", "8"), v["NewPassword"][0].Message)
}

func Test_DTO_resetPassword_notOk_2(t *testing.T) {
	u := ResetPassword{
		Token:       "098sdf",
		NewPassword: "testtest",
	}
	err := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "ResetPassword"), v["NewPassword"][0].Message)
}

func Test_DTO_resetPassword_ok(t *testing.T) {
	u := ResetPassword{
		Token:       "098sdf",
		NewPassword: "testTEST123*",
	}
	err := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	assert.Nil(t, err)
}
