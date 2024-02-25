package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
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
	errs := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "NewPassword"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_resetPassword_ok(t *testing.T) {
	u := ResetPassword{
		Token:       "098sdf",
		NewPassword: "testTEST123*",
	}
	errs := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	assert.Nil(t, errs)
}
