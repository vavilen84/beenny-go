package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"testing"
	"time"
)

func Test_Unit_ScenarioCreate_notOk_required(t *testing.T) {
	u := JWTInfo{}
	errs := validation.ValidateByScenario(constants.ScenarioCreate, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "UserId"),
		fmt.Sprintf(constants.RequiredErrorMsg, "Secret"),
		fmt.Sprintf(constants.RequiredErrorMsg, "ExpiresAt"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}

func Test_Unit_ScenarioCreate_notOk_expiredInPast(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	u := JWTInfo{
		UserId:    15,
		Secret:    "123123123",
		ExpiresAt: past,
	}
	errs := validation.ValidateByScenario(constants.ScenarioCreate, u)
	mustHaveErrors := []string{
		fmt.Sprintf(constants.FutureErrorMsg, "ExpiresAt"),
	}
	ok := helpers.AllErrorsExist(mustHaveErrors, errs)
	assert.True(t, ok)
}
