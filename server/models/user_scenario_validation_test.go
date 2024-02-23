package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_User_ScenarioCreate_notOk(t *testing.T) {
	u := User{}
	errs := validation.ValidateByScenario(constants.ScenarioCreate, u)
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
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioCreate_ok(t *testing.T) {
	u := GetTestValidUserModel()
	errs := validation.ValidateByScenario(constants.ScenarioCreate, u)
	assert.Nil(t, errs)
}

func Test_User_ScenarioHashPassword_notOk(t *testing.T) {
	u := User{}
	errs := validation.ValidateByScenario(constants.ScenarioHashPassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
		fmt.Sprintf(constants.RequiredErrorMsg, "PasswordSalt"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioHashPassword_Ok(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	u.EncodePassword()
	errs := validation.ValidateByScenario(constants.ScenarioHashPassword, u)
	assert.Nil(t, errs)
}

func Test_User_ScenarioForgotPassword_notOk(t *testing.T) {
	u := User{}
	errs := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "PasswordResetToken"),
		fmt.Sprintf(constants.RequiredErrorMsg, "PasswordResetTokenExpiresAt"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioForgotPassword_Ok(t *testing.T) {
	u := User{}
	u.SetForgotPasswordData()
	errs := validation.ValidateByScenario(constants.ScenarioForgotPassword, u)
	assert.Nil(t, errs)
}

func Test_User_ScenarioChangePassword_notOk(t *testing.T) {
	u := User{}
	errs := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioChangePassword_Ok(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	errs := validation.ValidateByScenario(constants.ScenarioChangePassword, u)
	assert.Nil(t, errs)
}

func Test_User_ScenarioResetPassword_notOk(t *testing.T) {
	u := User{}
	errs := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Password"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioResetPassword_Ok(t *testing.T) {
	u := User{
		Password: "12345678lT*",
	}
	errs := validation.ValidateByScenario(constants.ScenarioResetPassword, u)
	assert.Nil(t, errs)
}

func Test_User_ScenarioVerifyEmail_notOk(t *testing.T) {
	u := User{
		IsEmailVerified: false,
		EmailTwoFaCode:  "123456",
	}
	errs := validation.ValidateByScenario(constants.ScenarioVerifyEmail, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.EqErrorMsg, "IsEmailVerified", "true"),
		fmt.Sprintf(constants.EqErrorMsg, "EmailTwoFaCode", ""),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioVerifyEmail_Ok(t *testing.T) {
	u := User{
		IsEmailVerified: true,
		EmailTwoFaCode:  "",
	}
	errs := validation.ValidateByScenario(constants.ScenarioVerifyEmail, u)
	assert.Nil(t, errs)
}

func Test_User_ScenarioLoginTwoFaStepOne_notOk(t *testing.T) {
	u := User{
		EmailTwoFaCode: "",
	}
	errs := validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_User_ScenarioLoginTwoFaStepOne_Ok(t *testing.T) {
	u := User{
		EmailTwoFaCode: "123456",
	}
	errs := validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, u)
	assert.Nil(t, errs)
}
