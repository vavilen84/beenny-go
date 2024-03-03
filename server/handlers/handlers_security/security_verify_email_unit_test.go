package handlers_security_test

import (
	"testing"
)

func Test_Unit_VerifyEmail_ok(t *testing.T) {
	//customMatcher := mocks.CustomMatcher{}
	//db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//gormDB := store.GetMockDB(db)
	//store.SetMockDb(gormDB)
	//
	//ts := makeTestServer()
	//defer ts.Close()
	//
	//// find user by email
	//// no user found
	//rows := sqlmock.NewRows([]string{"id"})
	//expectedSQL := "SELECT * FROM `users`"
	//sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)
	//
	//// insert new user
	//sql := "INSERT INTO `users`"
	//sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))
	//
	//// send verification email
	//mockSESclient := mocks.NewSESClient(t)
	//aws.SetSESClient(mockSESclient)
	//mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	//
	//registerUser(t, ts)
	//
	//mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	//mockSESclient.AssertExpectations(t)
	//
	//if err := sqlMock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}
}
