package handlers_security_test

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/test"
	"testing"
)

func Test_Integration_Security_Register_OK(t *testing.T) {
	// TODO: need to add unit tests first
	t.Skip()
	ts := test.InitTestApp()
	defer ts.Close()
	u := test.RegisterUser(t, ts)
	db := store.GetDB()

	requestModel := test.GetValidRegisterUserDTO()

	userFromDb, err := models.FindUserById(db, u.Id)
	assert.Nil(t, err)

	assert.NotEmpty(t, u.PasswordSalt)
	passwordIsValid := password.Verify(requestModel.Password, userFromDb.PasswordSalt, userFromDb.Password, nil)
	assert.True(t, passwordIsValid)
}
