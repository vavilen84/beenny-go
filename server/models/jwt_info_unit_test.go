package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/store"
	"testing"
	"time"
)

func Test_Unit_InsertJWTInfo_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	sqlMock.ExpectExec("INSERT INTO `jwt_info`").WillReturnResult(sqlmock.NewResult(1, 1))

	future := time.Now().Add(1 * time.Hour)
	m := JWTInfo{
		UserId:    15,
		Secret:    "123123123",
		ExpiresAt: future,
	}

	err = InsertJWTInfo(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_FindJWTInfoById_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	expectedSQL := "SELECT * FROM `jwt_info`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	mm, err := FindJWTInfoById(gormDB, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, mm.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_GenerateSecret_ok(t *testing.T) {
	m := JWTInfo{}
	m.GenerateSecret()
	assert.NotEmpty(t, m.Secret)
}
