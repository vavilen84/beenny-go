package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_DTO_TwoFaLoginStepOne_notOk_1(t *testing.T) {
	u := TwoFaLoginStepOne{
		Email:    "",
		Password: "",
	}
	errs := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepOne, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Email"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_TwoFaLoginStepOne_notOk_2(t *testing.T) {
	u := TwoFaLoginStepOne{
		Email:    "not_valid_email",
		Password: "not_valid_pass",
	}
	errs := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepOne, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.EmailErrorMsg),
		fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "Password"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_TwoFaLoginStepOne_ok(t *testing.T) {
	u := TwoFaLoginStepOne{
		Email:    "user@example.com",
		Password: "testTEST123*",
	}
	errs := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepOne, u)
	assert.Nil(t, errs)
}
