package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/auth"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/test"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
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

func verifyEmail(t *testing.T, ts *httptest.Server, u models.User) {
	//url := ts.URL + "/api/v1/security/verify-email?token=" + u.EmailTwoFaCode
	//req, err := http.NewRequest(http.MethodGet, url, nil)
	//if err != nil {
	//	t.Fatalf("Failed to create request: %v", err)
	//}
	//
	//res, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	t.Fatalf("Failed to send request: %v", err)
	//}
	//defer res.Body.Close()
	//
	//responseBody, err := io.ReadAll(res.Body)
	//if err != nil {
	//	t.Fatalf("Error reading response body: %v", err)
	//}
	//
	//registerResp := dto.Response{}
	//err = json.Unmarshal(responseBody, &registerResp)
	//if err != nil {
	//	t.Fatalf("Error reading response body: %v", err)
	//}
	//
	//if res.StatusCode != http.StatusOK {
	//	t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	//}
	//
	//assert.Equal(t, registerResp.Status, http.StatusOK)
	//assert.Empty(t, registerResp.Error)
	//assert.Empty(t, registerResp.Errors)
	//assert.Empty(t, registerResp.FormErrors)
}

func loginUser(t *testing.T, ts *httptest.Server) string {
	twoFaLoginFirstStep(t, ts)
	jwtTok := twoFaLoginSecondStep(t, ts)
	return jwtTok
}

func performRequestBase(req *http.Request, t *testing.T, authToken *string) ([]byte, int) {
	if authToken != nil {
		req.Header.Add("Authorization", "Bearer "+*authToken)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	return body, resp.StatusCode
}

func performRequest(req *http.Request, t *testing.T, authToken *string) ([]byte, int) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return performRequestBase(req, t, authToken)
}

func deleteRequest(t *testing.T, url string, authToken *string) ([]byte, int) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	return performRequest(req, t, authToken)
}

func post(t *testing.T, url string, byteBody []byte, authToken *string) ([]byte, int) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	return performRequest(req, t, authToken)
}

func put(t *testing.T, url string, byteBody []byte, authToken *string) ([]byte, int) {
	req, err := http.NewRequest("PUT", url, bytes.NewReader(byteBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	return performRequest(req, t, authToken)
}

func get(t *testing.T, url string, authToken *string) ([]byte, int) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	return performRequest(req, t, authToken)
}

func twoFaLoginSecondStep(t *testing.T, ts *httptest.Server) (jwtToken string) {
	db := store.GetDB()
	u, err := models.FindUserByEmail(db, test.TestUserEmail)
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
		Email:    test.TestUserEmail,
		Password: test.TestUserPassword,
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

func registerUser(t *testing.T, userInput *dto.Register) models.User {
	i := dto.Register{
		Password: UserPassword,
	}
	if userInput != nil {
		i = *userInput
	} else {
		i = getBaseRegisterInput()
	}

	byteBody, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	body, statusCode := post(t, "/api/v1/security/register", byteBody, nil)

	newCreatedUser := models.User{}
	err = json.Unmarshal(body, &newCreatedUser)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, i.Email, newCreatedUser.Email)
	assert.NotEmpty(t, newCreatedUser.Id)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, statusCode)
	}

	return newCreatedUser
}

func marshalBodyGeneric[T any](m T) []byte {
	byteBody, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return byteBody
}

func unmarshalBodyGeneric[T any](body []byte, m T) T {
	err := json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}

	return m
}
