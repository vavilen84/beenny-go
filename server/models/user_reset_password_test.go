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

func TestUser_ResetPassword_passwordValidation(t *testing.T) {
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

	// should be error
	m := User{
		Password: "1234567",
	}
	err = UserResetPassword(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "User", "Password", "8"), v["Password"][0].Message)

	// should be error
	m = User{
		Password: "1234567+",
	}
	err = UserResetPassword(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "User"), v["Password"][0].Message)

	mock.ExpectExec("UPDATE user").WillReturnResult(sqlmock.NewResult(1, 1))

	// no error
	rowValidPassword := "12345678lT*"
	m.Password = rowValidPassword
	err = UserResetPassword(gormDB, &m)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
