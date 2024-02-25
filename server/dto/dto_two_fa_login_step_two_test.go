package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_DTO_TwoFaLoginStepTwo_notOk_1(t *testing.T) {
	u := TwoFaLoginStepTwo{
		EmailTwoFaCode: "",
	}
	errs := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_TwoFaLoginStepTwo_notOk_2(t *testing.T) {
	u := TwoFaLoginStepTwo{
		EmailTwoFaCode: "1234567",
	}
	errs := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.MaxValueErrorMsg, "EmailTwoFaCode"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_DTO_TwoFaLoginStepTwo_ok(t *testing.T) {
	u := TwoFaLoginStepTwo{
		EmailTwoFaCode: "123456",
	}
	errs := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, u)
	assert.Nil(t, errs)
}
