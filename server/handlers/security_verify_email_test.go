package handlers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"log"
	"net/http"
	"testing"
)

func TestVerifyEmail_OK(t *testing.T) {

	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)
	db := store.GetDB()

	u, err := models.FindUserByEmail(db, registerUserEmail)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotEmpty(t, u.EmailTwoFaCode)
	assert.Empty(t, u.IsEmailVerified)

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(ts.URL+"/api/v1/security/verify-email?token=%s", u.EmailTwoFaCode),
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	u, err = models.FindUserByEmail(db, registerUserEmail)
	if err != nil {
		log.Fatal(err)
	}
	assert.Empty(t, u.EmailTwoFaCode)
	assert.NotEmpty(t, u.IsEmailVerified)

}
