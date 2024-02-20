package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"strings"
	"testing"
)

import "errors"

type CustomMatcher struct{}

func (c CustomMatcher) Match(expectedSQL, actualSQL string) error {
	if !strings.Contains(actualSQL, expectedSQL) {
		return errors.New("SQL doesnt match")
	}
	return nil
}

func Test_User_ScenarioCreate_notOk(t *testing.T) {
	u := User{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	fmt.Printf("%v", v)
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
	u := User{
		FirstName:      "John",
		LastName:       "Dou",
		Email:          "email@example.com",
		CurrentCountry: "UA",
		CountryOfBirth: "UA",
		Gender:         constants.GenderMale,
		Timezone:       "US/Arizona",
		Birthday:       "1984-01-23",
		Password:       "12345678lT*",
		Photo:          "/2024/01/23/s09d8fs09dfu.jpg",
		Role:           constants.RoleUser,
		EmailTwoFaCode: "123456",
	}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	assert.Nil(t, err)
}