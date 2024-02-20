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
	"time"
)

func TestUser_ForgotPassword_passwordValidation(t *testing.T) {
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

	// Get the current time
	currentTime := time.Now()

	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}

	// Subtract 24 hours from the current time
	pastTime := currentTime.Add(duration)

	// should be error
	m := User{
		PasswordResetToken:         "",
		PasswordResetTokenExpireAt: &pastTime,
	}
	err = ForgotPassword(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "User", "PasswordResetToken", "3"), v["PasswordResetToken"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.FutureErrorMsg, "User", "PasswordResetTokenExpireAt"), v["PasswordResetTokenExpireAt"][0].Message)

	// should be error
	m = User{
		PasswordResetTokenExpireAt: nil,
	}
	err = ForgotPassword(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "User", "PasswordResetTokenExpireAt"), v["PasswordResetTokenExpireAt"][0].Message)

	// Calculate the duration of 24 hours
	duration, err = time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime = time.Now()
	// Add 24 hours to the current time
	futureTime := currentTime.Add(duration)

	mock.ExpectExec("UPDATE user").WillReturnResult(sqlmock.NewResult(1, 1))

	// no error
	m = User{
		PasswordResetToken:         "token",
		PasswordResetTokenExpireAt: &futureTime,
	}
	err = ForgotPassword(gormDB, &m)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
