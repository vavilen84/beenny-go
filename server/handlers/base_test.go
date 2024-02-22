package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/auth"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/handlers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	// user is SES verified
	registerUserEmail    = "vladimir.teplov@gmail.com"
	registerUserPassword = "testTEST123*"
)

type TestRegisterResp struct {
	Status     int                 `json:"status"`
	Data       interface{}         `json:"data"`
	Error      string              `json:"error"`
	Errors     map[string][]string `json:"errors"`
	FormErrors map[string][]string `json:"formErrors"`
}

type TestTwoFaLoginFirstResp struct {
	Status     int                 `json:"status"`
	Data       interface{}         `json:"data"`
	Error      string              `json:"error"`
	Errors     map[string][]string `json:"errors"`
	FormErrors map[string][]string `json:"formErrors"`
}

type TestTwoFaLoginSecondRespDataToken struct {
	Token string `json:"token"`
}

type TestTwoFaLoginSecondResp struct {
	Status     int                               `json:"status"`
	Data       TestTwoFaLoginSecondRespDataToken `json:"data"`
	Error      string                            `json:"error"`
	Errors     map[string][]string               `json:"errors"`
	FormErrors map[string][]string               `json:"formErrors"`
}

func makeTestServer() *httptest.Server {
	handler := handlers.MakeHandler()
	ts := httptest.NewServer(handler)
	return ts
}

func initApp() *httptest.Server {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Can not load .env file")
	}
	store.InitTestDB()
	db := store.GetDB()
	if err := db.Exec("delete from `jwt_info`").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	if err := db.Exec("delete from `users`").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	return makeTestServer()
}

func getValidRegisterUserDTO() dto.Register {
	return dto.Register{
		FirstName:       "John",
		LastName:        "Dou",
		CurrentCountry:  "UA",
		CountryOfBirth:  "UA",
		Gender:          constants.GenderMale,
		Timezone:        "US/Arizona",
		Birthday:        "1984-01-23",
		Photo:           "/2024/01/23/s09d8fs09dfu.jpg",
		Email:           registerUserEmail,
		Password:        registerUserPassword,
		ConfirmPassword: registerUserPassword,
		AgreeTerms:      true,
	}
}

func verifyEmail(t *testing.T, ts *httptest.Server, u models.User) {
	url := ts.URL + "/api/v1/security/verify-email?token=" + u.EmailTwoFaCode
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	registerResp := dto.Response{}
	err = json.Unmarshal(responseBody, &registerResp)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	assert.Equal(t, registerResp.Status, http.StatusOK)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Errors)
	assert.Empty(t, registerResp.FormErrors)
}

func registerUser(t *testing.T, ts *httptest.Server) models.User {
	body := getValidRegisterUserDTO()
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

	assert.Equal(t, responseBodyDataWrapper.Status, http.StatusOK)
	assert.Empty(t, responseBodyDataWrapper.Error)
	assert.Empty(t, responseBodyDataWrapper.Errors)
	assert.Empty(t, responseBodyDataWrapper.FormErrors)

	responseBodyData, ok := responseBodyDataWrapper.Data.(map[string]interface{})
	assert.True(t, ok)
	u := models.User{}
	err = json.Unmarshal(responseBodyData["user"].([]byte), &u)
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

	assert.NotEmpty(t, u.PasswordSalt)
	assert.NotEmpty(t, u.Id)

	assert.False(t, u.IsEmailVerified)

	passwordIsValid := password.Verify(body.Password, u.PasswordSalt, u.Password, nil)
	assert.True(t, passwordIsValid)

	return u
}

func loginUser(t *testing.T, ts *httptest.Server) string {
	twoFaLoginFirstStep(t, ts)
	jwtTok := twoFaLoginSecondStep(t, ts)
	return jwtTok
}

func twoFaLoginSecondStep(t *testing.T, ts *httptest.Server) (jwtToken string) {
	db := store.GetDB()
	u, err := models.FindUserByEmail(db, registerUserEmail)
	if err != nil {
		log.Fatal(err)
	}

	body := dto.TwoFaLoginStepTwo{
		EmailTwoFaCode: u.EmailTwoFaCode,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/two-fa-login-step-two", bytes.NewReader(bodyBytes))
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
		log.Fatal(err)
	}

	twoFaLoginSecondStep := TestTwoFaLoginSecondResp{}
	err = json.Unmarshal(responseBody, &twoFaLoginSecondStep)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	assert.Equal(t, twoFaLoginSecondStep.Status, http.StatusOK)
	assert.NotEmpty(t, twoFaLoginSecondStep.Data.Token)
	assert.Empty(t, twoFaLoginSecondStep.Error)
	assert.Empty(t, twoFaLoginSecondStep.Error)
	assert.Empty(t, twoFaLoginSecondStep.Errors)
	assert.Empty(t, twoFaLoginSecondStep.FormErrors)

	return twoFaLoginSecondStep.Data.Token
}

func twoFaLoginFirstStep(t *testing.T, ts *httptest.Server) {

	body := dto.TwoFaLoginStepOne{
		Email:    registerUserEmail,
		Password: registerUserPassword,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/two-fa-login-step-one", bytes.NewReader(bodyBytes))
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
		log.Fatal(err)
	}

	twoFaLoginStepFirst := TestTwoFaLoginFirstResp{}
	err = json.Unmarshal(responseBody, &twoFaLoginStepFirst)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	assert.Equal(t, twoFaLoginStepFirst.Status, http.StatusOK)
	assert.Empty(t, twoFaLoginStepFirst.Data)
	assert.Empty(t, twoFaLoginStepFirst.Error)
	assert.Empty(t, twoFaLoginStepFirst.Error)
	assert.Empty(t, twoFaLoginStepFirst.Errors)
	assert.Empty(t, twoFaLoginStepFirst.FormErrors)
}

func checkToken(t *testing.T, db *gorm.DB, token string) *models.User {
	isValid, err := auth.VerifyJWT(db, []byte(token))
	if err != nil || token == "" || !isValid {
		log.Fatalln(err)
	}

	jwtPayload, err := auth.ParseJWTPayload([]byte(token))
	if err != nil {
		log.Fatalln(err)
	}
	assert.NotEmpty(t, jwtPayload.JWTInfoId)

	jwtInfo, err := models.FindJWTInfoById(db, jwtPayload.JWTInfoId)
	if err != nil {
		log.Fatalln(err)
	}

	userByJWTInfo, err := models.FindUserById(db, jwtInfo.UserId)
	if err != nil {
		log.Fatalln(err)
	}

	return userByJWTInfo
}
