package handlers_security_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stretchr/testify/mock"
	"github.com/vavilen84/beenny-go/aws"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/test"
	"testing"
)

func Test_Unit_ChangePassword_ok(t *testing.T) {
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
