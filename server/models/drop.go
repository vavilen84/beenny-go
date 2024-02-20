package models

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"time"
)

type Drop struct {
	Id                   int         `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	CollectionName       string      `json:"collection_name"  `
	Blockchain           Blockchain  `json:"blockchain" column:"blockchain"`
	BlockchainName       string      `json:"blockchain_name" column:"blockchain_name"`
	WebsiteURL           string      `json:"website_url" column:"website_url"`
	TwitterURL           string      `json:"twitter_url" column:"twitter_url"`
	DiscordURL           string      `json:"discord_url" column:"discord_url"`
	PublicSaleDateTime   time.Time   `json:"public_sale_date" column:"public_sale_date"`
	TimeZone             string      `json:"time_zone" column:"time_zone"`
	PublicSalePrice      float64     `json:"public_sale_price" column:"public_sale_price"`
	TotalSupply          int         `json:"total_supply" column:"total_supply"`
	UserID               int         `json:"user_id" column:"user_id"`
	BillingPlan          BillingPlan `json:"billing_plan" column:"billing_plan"`
	BillingTransactionID string      `json:"billing_transaction_id" column:"billing_transaction_id"`
	Status               DropStatus  `json:"status" column:"status"`
	PreviewImg           string      `json:"preview_img" column:"preview_img"`
	BannerImg            string      `json:"banner_img" column:"banner_img"`
}

func (m *Drop) TableName() string {
	return "drop"
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
			"UserID":             "required",
			"BillingPlan":        "required,gt=0,lt=3",
			"Status":             "required,gt=0,lt=3",
		},
	}
}

func InsertDrop(db *gorm.DB, m *Drop) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}

	if m.WebsiteURL == "" && m.DiscordURL == "" && m.TwitterURL == "" {
		err = errors.New(AtLeastOneWebsiteOrGroupLinkErrMsg)
		helpers.LogError(err)
		return
	}

	if m.Blockchain == OtherBlockchain && m.BlockchainName == "" {
		err = errors.New(BlockchainNameRequiredErrMsg)
		helpers.LogError(err)
		return
	}

	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func UpdateDrop(db *gorm.DB, m *Drop) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}

	if m.WebsiteURL == "" && m.DiscordURL == "" && m.TwitterURL == "" {
		err = errors.New(AtLeastOneWebsiteOrGroupLinkErrMsg)
		helpers.LogError(err)
		return
	}

	if m.Blockchain == OtherBlockchain && m.BlockchainName == "" {
		err = errors.New(BlockchainNameRequiredErrMsg)
		helpers.LogError(err)
		return
	}

	err = db.Save(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func FindDropById(db *gorm.DB, id int) (*Drop, error) {
	m := Drop{}
	err := db.Where("id = ?", id).First(&m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}
