package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"log"
	"testing"
)

func Test_DTO_resetPassword_notOk_1(t *testing.T) {
	u := ResetPassword{
		Token:       "",
		NewPassword: "",
	}
	errs := validation.ValidateByScenario(constants.ScenarioResetPassword, u)

	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Token"),
		fmt.Sprintf(constants.RequiredErrorMsg, "NewPassword"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
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
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg), v["NewPassword"][0].Message)
}

func Test_DTO_resetPassword_ok(t *testing.T) {
	u := ResetPassword{
		Token:       "098sdf",
		NewPassword: "testTEST123*",
	}
	err := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	assert.Nil(t, err)
}
