package models

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func TestInsertJWTInfo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	// error
	m := JWTInfo{}
	err = InsertJWTInfo(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "JWTInfo", "UserId"), v["UserId"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "JWTInfo", "ExpiresAt"), v["ExpiresAt"][0].Message)

	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	pastTime := currentTime.Add(duration)

	// error
	m = JWTInfo{
		ExpiresAt: pastTime,
	}
	err = InsertJWTInfo(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.FutureErrorMsg, "JWTInfo", "ExpiresAt"), v["ExpiresAt"][0].Message)

	// Calculate the duration of 24 hours
	duration, err = time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime = time.Now()
	// Add 24 hours to the current time
	futureTime := currentTime.Add(duration)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `jwt_info`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// no error
	m = JWTInfo{
		UserId:    1,
		ExpiresAt: futureTime,
	}
	err = InsertJWTInfo(gormDB, &m)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_FindJWTInfoById(t *testing.T) {
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

	columns := []string{"id"}
	mock.ExpectQuery("SELECT * FROM `jwt_info`").
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	_, err = FindJWTInfoById(gormDB, 1)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
