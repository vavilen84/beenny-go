package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/validation"
	"log"
	"testing"
	"time"
)

func Test_DTO_DropCreate_notOk(t *testing.T) {

	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	pastTime := currentTime.Add(duration)

	u := Drop{
		CollectionName:     "",
		Blockchain:         0,
		PublicSaleDateTime: pastTime,
		TimeZone:           "",
		PublicSalePrice:    0,
		TotalSupply:        0,
		BillingPlan:        0,
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "CollectionName"), v["CollectionName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "Blockchain"), v["Blockchain"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.FutureErrorMsg, "Drop", "PublicSaleDateTime"), v["PublicSaleDateTime"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "TimeZone"), v["TimeZone"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "PublicSalePrice"), v["PublicSalePrice"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "TotalSupply"), v["TotalSupply"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "BillingPlan"), v["BillingPlan"][0].Message)
}

func Test_DTO_DropCreate_Ok(t *testing.T) {

	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	futureTime := currentTime.Add(duration)

	u := Drop{
		CollectionName:     "Col #1",
		Blockchain:         int(models.EthereumBlockchain),
		PublicSaleDateTime: futureTime,
		TimeZone:           "America/Adak",
		PublicSalePrice:    1,
		TotalSupply:        5000,
		BillingPlan:        int(models.StandardBillingPlan),
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, u)
	assert.Nil(t, err)
}

func Test_DTO_DropUpdate_notOk(t *testing.T) {

	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	pastTime := currentTime.Add(duration)

	u := Drop{
		CollectionName:     "",
		Blockchain:         0,
		PublicSaleDateTime: pastTime,
		TimeZone:           "",
		PublicSalePrice:    0,
		TotalSupply:        0,
		BillingPlan:        0,
	}
	err = validation.ValidateByScenario(constants.ScenarioUpdate, u)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "Id"), v["Id"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "CollectionName"), v["CollectionName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "Blockchain"), v["Blockchain"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.FutureErrorMsg, "Drop", "PublicSaleDateTime"), v["PublicSaleDateTime"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "TimeZone"), v["TimeZone"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "PublicSalePrice"), v["PublicSalePrice"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "TotalSupply"), v["TotalSupply"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "BillingPlan"), v["BillingPlan"][0].Message)
}

func Test_DTO_DropUpdate_Ok(t *testing.T) {

	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	futureTime := currentTime.Add(duration)

	u := Drop{
		Id:                 1,
		CollectionName:     "Col #1",
		Blockchain:         int(models.EthereumBlockchain),
		PublicSaleDateTime: futureTime,
		TimeZone:           "America/Adak",
		PublicSalePrice:    1,
		TotalSupply:        5000,
		BillingPlan:        int(models.StandardBillingPlan),
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, u)
	assert.Nil(t, err)
}
