package models

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"time"
)

type JWTInfo struct {
	Id        int       `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId    int       `column:"user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	Secret    string    `column:"secret"`
	CreatedAt time.Time `json:"created_at" column:"created_at" gorm:"autoCreateTime"`
	ExpiresAt time.Time `json:"expires_at" column:"expires_at"`
}

func (m *JWTInfo) TableName() string {
	return "jwt_info"
}

func (JWTInfo) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"UserId":    "required",
			"Secret":    "min=3,max=5000,required",
			"ExpiresAt": "required,customFutureValidator",
		},
	}
}

func (JWTInfo) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customFutureValidator", validation.CustomFutureValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func InsertJWTInfo(db *gorm.DB, m *JWTInfo) (err error) {
	m.generateSecret()
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func (m *JWTInfo) generateSecret() {
	m.Secret = helpers.GenerateRandomString(64)
}

func FindJWTInfoById(db *gorm.DB, id int) (*JWTInfo, error) {
	m := JWTInfo{}
	err := db.First(&m, id).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}
