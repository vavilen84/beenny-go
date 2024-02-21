package models

import (
	"database/sql"
	"errors"
	"github.com/vavilen84/nft-project/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type CustomMatcher struct{}

func (c CustomMatcher) Match(expectedSQL, actualSQL string) error {
	if !strings.Contains(actualSQL, expectedSQL) {
		return errors.New("SQL doesnt match")
	}
	return nil
}

func GetMockDB(db *sql.DB) (gormDB *gorm.DB) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB = gormDB.Session(&gorm.Session{SkipDefaultTransaction: true})

	return
}

func GetTestValidUserModel() User {
	return User{
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
}
