package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"log"
	"testing"
)

func Test_User_ScenarioCreate_notOk(t *testing.T) {
	u := User{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.EmailErrorMsg),
		fmt.Sprintf(constants.RequiredErrorMsg, "FirstName"),
		fmt.Sprintf(constants.RequiredErrorMsg, "LastName"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Gender"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Timezone"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Birthday"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Photo"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
		fmt.Sprintf(constants.RequiredErrorMsg, "CurrentCountry"),
		fmt.Sprintf(constants.RequiredErrorMsg, "CountryOfBirth"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Role"),
		fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioCreate_ok(t *testing.T) {
	u := GetTestValidUserModel()
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	assert.Nil(t, err)
}

func Test_User_ScenarioHashPassword_notOk(t *testing.T) {
	u := User{}
	err := validation.ValidateByScenario(constants.ScenarioHashPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
		fmt.Sprintf(constants.RequiredErrorMsg, "PasswordSalt"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioHashPassword_Ok(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	u.EncodePassword()
	err := validation.ValidateByScenario(constants.ScenarioHashPassword, u)
	assert.Nil(t, err)
}

func Test_User_ScenarioForgotPassword_notOk(t *testing.T) {
	u := User{}
	err := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "PasswordResetToken"),
		fmt.Sprintf(constants.RequiredErrorMsg, "PasswordResetTokenExpiresAt"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioForgotPassword_Ok(t *testing.T) {
	u := User{}
	u.SetForgotPasswordData()
	err := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	assert.Nil(t, err)
}

func Test_User_ScenarioChangePassword_notOk(t *testing.T) {
	u := User{}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioChangePassword_Ok(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	err := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	assert.Nil(t, err)
}

func Test_User_ScenarioResetPassword_notOk(t *testing.T) {
	u := User{}
	err := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioResetPassword_Ok(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	err := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	assert.Nil(t, err)
}

func Test_User_ScenarioVerifyEmail_notOk(t *testing.T) {
	u := User{
		IsEmailVerified: false,
		EmailTwoFaCode:  "123456",
	}
	err := validation.ValidateByScenario(constants.ScenarioVerifyEmail, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.EqErrorMsg, "IsEmailVerified", "true"),
		fmt.Sprintf(constants.EqErrorMsg, "EmailTwoFaCode", ""),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioVerifyEmail_Ok(t *testing.T) {
	u := User{
		IsEmailVerified: true,
		EmailTwoFaCode:  "",
	}
	err := validation.ValidateByScenario(constants.ScenarioVerifyEmail, u)
	assert.Nil(t, err)
}

func Test_User_ScenarioLoginTwoFaStepOne_notOk(t *testing.T) {
	u := User{
		EmailTwoFaCode: "",
	}
	err := validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}

func Test_User_ScenarioLoginTwoFaStepOne_Ok(t *testing.T) {
	u := User{
		EmailTwoFaCode: "123456",
	}
	err := validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, u)
	assert.Nil(t, err)
}
