package models

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestUser_SetUserEmailVerified(t *testing.T) {
	customMatcher := CustomMatcher{}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	mock.ExpectExec("UPDATE user").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// should be error
	m := User{
		EmailTwoFaCode: "somevalue",
	}
	err = SetUserEmailVerified(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.EqErrorMsg, "User", "IsEmailVerified", "true"), v["IsEmailVerified"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.EqErrorMsg, "User", "EmailTwoFaCode", ""), v["EmailTwoFaCode"][0].Message)

	// no error

	m = User{
		IsEmailVerified: true,
		EmailTwoFaCode:  "",
	}
	err = SetUserEmailVerified(gormDB, &m)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
