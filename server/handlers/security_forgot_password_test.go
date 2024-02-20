package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"log"
	"net/http"
	"testing"
)

func Test_ForgotPassword_userNotFound(t *testing.T) {
	ts := initApp()
	defer ts.Close()

	body := dto.ForgotPassword{
		Email: "not_existing_user@example.com",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/forgot-password", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, res.StatusCode)
	}
}

func Test_ForgotPassword_OK(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)

	db := store.GetDB()

	u, err := models.FindUserByEmail(db, registerUserEmail)
	if err != nil {
		log.Fatal("user not found")
	}
	assert.Empty(t, u.PasswordResetTokenExpireAt)
	assert.Empty(t, u.PasswordResetToken)

	body := dto.ForgotPassword{
		Email: registerUserEmail,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/forgot-password", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	u, err = models.FindUserByEmail(db, registerUserEmail)
	if err != nil {
		log.Fatal("user not found")
	}
	assert.NotEmpty(t, u.PasswordResetTokenExpireAt)
	assert.NotEmpty(t, u.PasswordResetToken)
}
