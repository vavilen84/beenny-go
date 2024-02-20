package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"time"
)

type Drop struct {
	Id                 int       `json:"id"`
	CollectionName     string    `json:"collection_name"`
	Blockchain         int       `json:"blockchain"`
	BlockchainName     string    `json:"blockchain_name"`
	WebsiteURL         string    `json:"website_url"`
	TwitterURL         string    `json:"twitter_url"`
	DiscordURL         string    `json:"discord_url"`
	PublicSaleDateTime time.Time `json:"public_sale_date"`
	TimeZone           string    `json:"time_zone"`
	PublicSalePrice    float64   `json:"public_sale_price"`
	TotalSupply        int       `json:"total_supply"`
	BillingPlan        int       `json:"billing_plan"`
}

func (Drop) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customFutureValidator", validation.CustomFutureValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func (Drop) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"CollectionName":     "required",
			"Blockchain":         "required,gt=0,lt=19",
			"PublicSaleDateTime": "required,customFutureValidator",
			"TimeZone":           "required",
			"PublicSalePrice":    "required",
			"TotalSupply":        "required",
			"BillingPlan":        "required,gt=0,lt=3",
		},
		constants.ScenarioUpdate: validation.FieldRules{
			"Id":                 "required",
			"CollectionName":     "required",
			"Blockchain":         "required,gt=0,lt=19",
			"PublicSaleDateTime": "required,customFutureValidator",
			"TimeZone":           "required",
			"PublicSalePrice":    "required",
			"TotalSupply":        "required",
			"BillingPlan":        "required,gt=0,lt=3",
		},
	}
}
