package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"

	"log"
	"testing"
)

func Test_Register_notOk_1(t *testing.T) {
	u := Register{
		Nickname:    "",
		Email:       "",
		Password:    "",
		BillingPlan: 10,
	}
	err := validation.ValidateByScenario(constants.ScenarioRegister, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "Register", "Nickname", "3"), v["Nickname"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "Register", "Email", "3"), v["Email"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "Register", "Password", "8"), v["Password"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.LowerThanTagErrorMsg, "Register", "4"), v["BillingPlan"][0].Message)
}

func Test_Register_notOk_2(t *testing.T) {
	u := Register{
		Nickname:    "nick",
		Email:       "not_valid_email",
		Password:    "12345678",
		BillingPlan: 0,
	}
	err := validation.ValidateByScenario(constants.ScenarioRegister, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg, "Register"), v["Email"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "Register"), v["Password"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Register", "BillingPlan"), v["BillingPlan"][0].Message)
}

func Test_Register_ok(t *testing.T) {
	u := Register{
		Nickname:    "nick",
		Email:       "email@example.com",
		Password:    "12345678lT*",
		BillingPlan: 1,
	}
	err := validation.ValidateByScenario(constants.ScenarioRegister, u)
	assert.Nil(t, err)
}
