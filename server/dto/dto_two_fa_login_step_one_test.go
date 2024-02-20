package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"testing"
)

func Test_DTO_TwoFaLoginStepOne_notOk_1(t *testing.T) {
	u := TwoFaLoginStepOne{
		Email:    "",
		Password: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepOne, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "TwoFaLoginStepOne", "Email", "3"), v["Email"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "TwoFaLoginStepOne", "Password", "8"), v["Password"][0].Message)
}

func Test_DTO_TwoFaLoginStepOne_notOk_2(t *testing.T) {
	u := TwoFaLoginStepOne{
		Email:    "not_valid_email",
		Password: "not_valid_pass",
	}
	err := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepOne, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg, "TwoFaLoginStepOne"), v["Email"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "TwoFaLoginStepOne"), v["Password"][0].Message)
}

func Test_DTO_TwoFaLoginStepOne_ok(t *testing.T) {
	u := TwoFaLoginStepOne{
		Email:    "user@example.com",
		Password: "testTEST123*",
	}
	err := validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepOne, u)
	assert.Nil(t, err)
}
