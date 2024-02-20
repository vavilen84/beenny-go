package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUser_FindUserById(t *testing.T) {
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
	mock.ExpectQuery("SELECT * FROM `user`").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	_, err = FindUserById(gormDB, 1)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_FindUserByTwoFAToken(t *testing.T) {
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

	columns := []string{"email_two_fa_code"}
	mock.ExpectQuery("SELECT * FROM `user`").
		WithArgs("token").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("token"))

	_, err = FindUserByTwoFAToken(gormDB, "token")
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_FindUserByEmail(t *testing.T) {
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

	columns := []string{"email"}
	mock.ExpectQuery("SELECT * FROM `user`").
		WithArgs("email@example.com").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("email@example.com"))

	_, err = FindUserByEmail(gormDB, "email@example.com")
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_FindUserByResetPasswordToken(t *testing.T) {
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

	columns := []string{"password_reset_token"}
	mock.ExpectQuery("SELECT * FROM `user`").
		WithArgs("token", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("token"))

	_, err = FindUserByResetPasswordToken(gormDB, "token")
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
