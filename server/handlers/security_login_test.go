package handlers_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/store"
	"testing"
)

func TestLogin_OK(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	db := store.GetDB()
	registerUser(t, ts)

	loggedInUserToken := loginUser(t, ts)
	loggedInUser := checkToken(t, db, loggedInUserToken)
	assert.Equal(t, loggedInUser.Email, registerUserEmail)
}
