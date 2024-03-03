package handlers_security_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vavilen84/beenny-go/aws"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/test"
	"log"
	"net/http"
	"testing"
)

func Test_Unit_Security_Register_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	// find user by email
	// no user found
	rows := sqlmock.NewRows([]string{"id"})
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	// insert new user
	sql := "INSERT INTO `users`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	// send verification email
	mockSESclient := mocks.NewSESClient(t)
	aws.SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)

	test.RegisterUser(t, ts)

	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Security_Register_notOk_can_not_parse_body(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	bodyBytes := make([]byte, 0)

	resp := test.Post(t, ts, constants.RegisterUserURL, bodyBytes, http.StatusBadRequest)
	assert.NotEmpty(t, resp.Errors)
	assert.Equal(t, constants.BadRequestError.Error(), resp.Errors[0])
	assert.Nil(t, resp.Data)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Security_Register_notOk_email_is_required(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	body := test.GetValidRegisterUserDTO()
	body.Email = "" // empty email
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	resp := test.Post(t, ts, constants.RegisterUserURL, bodyBytes, http.StatusBadRequest)
	assert.NotEmpty(t, resp.Errors)
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg), resp.Errors[0])
	assert.Nil(t, resp.Data)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Security_Register_notOk_passwords_dont_match(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	body := test.GetValidRegisterUserDTO()
	body.Password = body.ConfirmPassword + "1" // passwords dont match
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	resp := test.Post(t, ts, constants.RegisterUserURL, bodyBytes, http.StatusBadRequest)
	assert.NotEmpty(t, resp.Errors)
	assert.Equal(t, constants.PasswordsDontMatch.Error(), resp.Errors[0])
	assert.Nil(t, resp.Data)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Security_Register_notOk_user_already_registered(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	body := test.GetValidRegisterUserDTO()
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	resp := test.Post(t, ts, constants.RegisterUserURL, bodyBytes, http.StatusBadRequest)
	assert.NotEmpty(t, resp.Errors)
	assert.Equal(t, fmt.Sprintf(constants.UserAlreadyRegisteredFormat, body.Email), resp.Errors[0])
	assert.Nil(t, resp.Data)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Security_Register_notOk_user_insert_error(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	rows := sqlmock.NewRows([]string{"id"})
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	err = errors.New("sql query error")
	sql := "INSERT INTO `users`"
	sqlMock.ExpectExec(sql).WillReturnError(err)

	body := test.GetValidRegisterUserDTO()
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	resp := test.Post(t, ts, constants.RegisterUserURL, bodyBytes, http.StatusInternalServerError)
	assert.NotEmpty(t, resp.Errors)
	assert.Equal(t, constants.ServerError.Error(), resp.Errors[0])
	assert.Nil(t, resp.Data)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Unit_Security_Register_notOk_SendEmailVerificationMail_error(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)
	store.SetMockDb(gormDB)

	ts := test.MakeTestServer()
	defer ts.Close()

	// find user by email
	// no user found
	rows := sqlmock.NewRows([]string{"id"})
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	// insert new user
	sql := "INSERT INTO `users`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	// send verification email
	mockSESclient := mocks.NewSESClient(t)
	aws.SetSESClient(mockSESclient)

	var sesErr awserr.Error
	sesErr = test.SesError{
		CodeData: ses.ErrCodeMessageRejected,
	}
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, sesErr)

	body := test.GetValidRegisterUserDTO()
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	resp := test.Post(t, ts, constants.RegisterUserURL, bodyBytes, http.StatusInternalServerError)
	assert.NotEmpty(t, resp.Errors)
	assert.Equal(t, constants.ServerError.Error(), resp.Errors[0])
	assert.Nil(t, resp.Data)

	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
