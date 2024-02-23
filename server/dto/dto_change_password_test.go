package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_PasswordValidation_notOk_required(t *testing.T) {
	u := ChangePassword{}
	errs := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "OldPassword"),
		fmt.Sprintf(constants.RequiredErrorMsg, "NewPassword"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_PasswordValidation_notOk_customValidation(t *testing.T) {
	u := ChangePassword{
		OldPassword: "12345678",
		NewPassword: "12345678",
	}
	errs := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "OldPassword"),
		fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "NewPassword"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_PasswordValidation_ok(t *testing.T) {
	u := ChangePassword{
		OldPassword: "12345678lT*",
		NewPassword: "12345678lT*",
	}
	errs := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	assert.Nil(t, errs)
}
