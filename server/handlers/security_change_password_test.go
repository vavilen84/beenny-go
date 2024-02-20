package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/vavilen84/nft-project/dto"
	"log"
	"net/http"
	"testing"
)

func Test_ChangePassword_NotAuthorized(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)

	body := dto.ChangePassword{
		OldPassword: "",
		NewPassword: "",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/change-password", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, res.StatusCode)
	}
}

func Test_ChangePassword_NotValidPassword(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)

	loggedInUserToken := loginUser(t, ts)

	body := dto.ChangePassword{
		OldPassword: "testTEST123!",
		NewPassword: "testTEST123*",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/change-password", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", loggedInUserToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, res.StatusCode)
	}
}

func Test_ChangePassword_OK(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)

	loggedInUserToken := loginUser(t, ts)

	body := dto.ChangePassword{
		OldPassword: registerUserPassword,
		NewPassword: "testTEST123!",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/change-password", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", loggedInUserToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}
}
