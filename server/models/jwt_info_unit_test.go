package models

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/validation"
	"log"
	"testing"
	"time"
)

func Test_Unit_InsertJWTInfo(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	// error
	m := JWTInfo{}
	err = InsertJWTInfo(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}

	mustHaveErrors := []string{
		fmt.Sprintf(constants.RequiredErrorMsg, "UserId"),
		fmt.Sprintf(constants.RequiredErrorMsg, "ExpiresAt"),
	}
	ok = helpers.AllErrorsExist(mustHaveErrors, v)
	assert.True(t, ok)

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
		UserId:    15,
		ExpiresAt: pastTime,
	}
	err = InsertJWTInfo(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	mustHaveErrors = []string{
		fmt.Sprintf(constants.FutureErrorMsg, "ExpiresAt"),
	}
	ok = helpers.AllErrorsExist(mustHaveErrors, v)
	assert.True(t, ok)

	// Calculate the duration of 24 hours
	duration, err = time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime = time.Now()
	// Add 24 hours to the current time
	futureTime := currentTime.Add(duration)

	sqlMock.ExpectExec("INSERT INTO `jwt_info`").WillReturnResult(sqlmock.NewResult(1, 1))

	// no error
	m = JWTInfo{
		UserId:    1,
		ExpiresAt: futureTime,
	}
	err = InsertJWTInfo(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_FindJWTInfoById(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	columns := []string{"id"}
	sqlMock.ExpectQuery("SELECT * FROM `jwt_info`").
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	_, err = FindJWTInfoById(gormDB, 1)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
