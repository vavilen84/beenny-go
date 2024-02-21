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

func TestUser_InsertUser_ok(t *testing.T) {
	customMatcher := CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := GetMockDB(db)

	sql := "INSERT INTO `users`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = InsertUser(gormDB, &m)
	//sqlMock.ExpectCommit()
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_nicknameValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
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
		//Nickname: "",
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "Nickname", "3"), v["Nickname"][0].Message)

	//m.Nickname = "valid.nickname"
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}

	// no error
	_, ok = v["Nickname"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_passwordValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
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
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "Password", "8"), v["Password"][0].Message)

	// should be error
	m = User{
		Password: "1234567+",
	}
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg), v["Password"][0].Message)

	// no error
	rowValidPassword := "12345678lT*"
	m.Password = rowValidPassword
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	_, ok = v["Password"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_roleValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
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
		Role: 0,
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Role"), v["Role"][0].Message)

	// should be error
	m = User{
		Role: 2,
	}
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.LowerThanTagErrorMsg, "2"), v["Role"][0].Message)

	m.Role = constants.RoleUser
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}

	// no error
	_, ok = v["Role"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_2FaCodeValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
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
		EmailTwoFaCode: "",
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "EmailTwoFaCode"), v["EmailTwoFaCode"][0].Message)

	m.EmailTwoFaCode = "123456"
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}

	// no error
	_, ok = v["EmailTwoFaCode"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_OkInsert(t *testing.T) {
	//db, mock, err := sqlmock.New()
	//if err != nil {
	//	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	//}
	//defer db.Close()
	//gormDB, err := gorm.Open(mysql.New(mysql.Config{
	//	SkipInitializeWithVersion: true,
	//	Conn:                      db,
	//}), &gorm.Config{})
	//
	//mock.ExpectBegin()
	//mock.ExpectExec("INSERT INTO `user`").WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectCommit()
	//
	//// no error
	//m := User{
	//	//Nickname:       "nick",
	//	Email:          "valid.email@example.com",
	//	Password:       "12345678lT*",
	//	Role:           constants.RoleUser,
	//	EmailTwoFaCode: "123456",
	//}
	//err = InsertUser(gormDB, &m)
	//assert.Nil(t, err)
	//
	//if err := mock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}
}
