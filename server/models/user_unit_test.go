package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/store"
	"testing"
	"time"
)

func TestUser_InsertUser_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "INSERT INTO `users`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = InsertUser(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ForgotPassword_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	m.SetForgotPasswordData()
	err = ForgotPassword(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_SetEmailTwoFaCode_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = SetEmailTwoFaCode(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ResetEmailTwoFaCode_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = ResetEmailTwoFaCode(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ResetResetPasswordToken_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = ResetResetPasswordToken(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_SetUserEmailVerified_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	m.IsEmailVerified = true
	m.EmailTwoFaCode = ""
	err = SetUserEmailVerified(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_UserResetPassword_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = UserResetPassword(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_UserChangePassword_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "UPDATE users"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

	m := GetTestValidUserModel()
	err = UserChangePassword(gormDB, &m)
	assert.Nil(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_encodePassword(t *testing.T) {
	m := GetTestValidUserModel()
	pass := m.Password
	m.encodePassword()
	assert.NotEqual(t, pass, m.Password)
	assert.NotEmpty(t, m.PasswordSalt)
}

func Test_FindUserById_ok(t *testing.T) {
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

	m, err := FindUserById(gormDB, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, m.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_FindUserByTwoFAToken_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	rows := sqlmock.NewRows([]string{"id", "email_two_fa_code"}).AddRow(1, "123456")
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	m, err := FindUserByTwoFAToken(gormDB, "123456")
	assert.Nil(t, err)
	assert.Equal(t, 1, m.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_FindUserByEmail_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	rows := sqlmock.NewRows([]string{"id", "email"}).AddRow(1, "user@example.com")
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	m, err := FindUserByEmail(gormDB, "user@example.com")
	assert.Nil(t, err)
	assert.Equal(t, 1, m.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_FindUserByResetPasswordToken_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	future := time.Now().Add(1 * time.Hour)
	rows := sqlmock.NewRows([]string{"id", "password_reset_token", "password_reset_token_expire_at"}).
		AddRow(1, "123456", &future)
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	m, err := FindUserByResetPasswordToken(gormDB, "123456")
	assert.Nil(t, err)
	assert.Equal(t, 1, m.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_SetForgotPasswordData_ok(t *testing.T) {
	u := User{}
	u.SetForgotPasswordData()
	assert.NotEmpty(t, u.PasswordResetToken)
	assert.NotEmpty(t, u.PasswordResetTokenExpiresAt)
}
