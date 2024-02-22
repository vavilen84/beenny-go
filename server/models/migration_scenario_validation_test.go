package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
)

func Test_Unit_Migration_ValidateOnCreate_notOk(t *testing.T) {
	u := Migration{}
	errs := validation.ValidateByScenario(constants.ScenarioCreate, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "Version"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Filename"),
		fmt.Sprintf(constants.RequiredErrorMsg, "CreatedAt"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_Unit_Migration_ValidateOnCreate_ok(t *testing.T) {
	u := Migration{
		Version:   1708526378,
		Filename:  "1708526378_add_users_table.up.sql",
		CreatedAt: 1708526378,
	}
	errs := validation.ValidateByScenario(constants.ScenarioCreate, u)
	assert.Empty(t, errs)
}
