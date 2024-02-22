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

func TestMigration_ValidateOnCreate(t *testing.T) {
	u := Migration{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Version"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Filename"),
		fmt.Sprintf(constants.RequiredErrorMsg, "CreatedAt"),
	}
	ok = helpers.AllStringsAreErrors(mustHaveErrors, v)
	assert.True(t, ok)
}
