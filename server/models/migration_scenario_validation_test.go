package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
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
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Version"), v["Version"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Filename"), v["Filename"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "CreatedAt"), v["CreatedAt"][0].Message)
}
