package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_DTO_ForgotPassword_notOk_required(t *testing.T) {
	u := ForgotPassword{}
	errs := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Email"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_ForgotPassword_notOk_email(t *testing.T) {
	u := ForgotPassword{
		Email: "not_valid_email",
	}
	errs := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.EmailErrorMsg),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_ForgotPassword_ok(t *testing.T) {
	u := ForgotPassword{
		Email: "user@example.com",
	}
	errs := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	assert.Nil(t, errs)
}
