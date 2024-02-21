package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/validation"
	"log"
	"testing"
)

func Test_DTO_TwoFaLoginStepTwo_notOk_1(t *testing.T) {
	u := TwoFaLoginStepTwo{
		EmailTwoFaCode: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "EmailTwoFaCode", "6"), v["EmailTwoFaCode"][0].Message)
}

func Test_DTO_TwoFaLoginStepTwo_notOk_2(t *testing.T) {
	u := TwoFaLoginStepTwo{
		EmailTwoFaCode: "1234567",
	}
	err := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MaxValueErrorMsg, "EmailTwoFaCode", "6"), v["EmailTwoFaCode"][0].Message)
}

func Test_DTO_TwoFaLoginStepTwo_ok(t *testing.T) {
	u := TwoFaLoginStepTwo{
		EmailTwoFaCode: "123456",
	}
	err := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, u)
	assert.Nil(t, err)
}
