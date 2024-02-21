package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/validation"
	"log"
	"testing"
)

func Test_DTO_ForgotPassword_notOk_1(t *testing.T) {
	u := ForgotPassword{
		Email: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "Email", "3"), v["Email"][0].Message)
}

func Test_DTO_ForgotPassword_notOk_2(t *testing.T) {
	u := ForgotPassword{
		Email: "not_valid_email",
	}
	err := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg), v["Email"][0].Message)
}

func Test_DTO_ForgotPassword_ok(t *testing.T) {
	u := ForgotPassword{
		Email: "user@example.com",
	}
	err := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	assert.Nil(t, err)
}
