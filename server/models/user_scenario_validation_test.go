package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
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
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg), v["Email"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "FirstName"), v["FirstName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "LastName"), v["LastName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Photo"), v["Photo"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Gender"), v["Gender"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Timezone"), v["Timezone"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Birthday"), v["Birthday"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Photo"), v["Photo"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Password"), v["Password"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "CurrentCountry"), v["CurrentCountry"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "CountryOfBirth"), v["CountryOfBirth"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Role"), v["Role"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"), v["EmailTwoFaCode"][0].Message)
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
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Password"), v["Password"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "PasswordSalt"), v["PasswordSalt"][0].Message)
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
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "PasswordResetToken"), v["PasswordResetToken"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "PasswordResetTokenExpiresAt"), v["PasswordResetTokenExpiresAt"][0].Message)
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
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Password"), v["Password"][0].Message)
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
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Password"), v["Password"][0].Message)
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
	assert.Equal(t, fmt.Sprintf(constants.EqErrorMsg, "IsEmailVerified", "true"), v["IsEmailVerified"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.EqErrorMsg, "EmailTwoFaCode", ""), v["EmailTwoFaCode"][0].Message)
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
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"), v["EmailTwoFaCode"][0].Message)
}

func Test_User_ScenarioLoginTwoFaStepOne_Ok(t *testing.T) {
	u := User{
		EmailTwoFaCode: "123456",
	}
	err := validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, u)
	assert.Nil(t, err)
}
