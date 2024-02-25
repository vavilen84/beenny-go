package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/models"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RegisterUser(t *testing.T, ts *httptest.Server) models.User {

	body := GetValidRegisterUserDTO()
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/register", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	responseBodyDataWrapper := dto.Response{}
	err = json.Unmarshal(responseBody, &responseBodyDataWrapper)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}
	assert.Empty(t, responseBodyDataWrapper.Errors)

	u := models.User{}
	err = json.Unmarshal(responseBodyDataWrapper.Data, &u)
	assert.Nil(t, err)

	assert.Equal(t, body.FirstName, u.FirstName)
	assert.Equal(t, body.LastName, u.LastName)
	assert.Equal(t, body.CurrentCountry, u.CurrentCountry)
	assert.Equal(t, body.CountryOfBirth, u.CountryOfBirth)
	assert.Equal(t, body.Gender, u.Gender)
	assert.Equal(t, body.Timezone, u.Timezone)
	assert.Equal(t, body.Birthday, u.Birthday)
	assert.Equal(t, body.Photo, u.Photo)
	assert.Equal(t, body.Email, u.Email)
	assert.NotEmpty(t, u.Id)
	assert.False(t, u.IsEmailVerified)

	return u
}
