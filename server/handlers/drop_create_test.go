package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/models"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestDropCreate_unauthorized(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)

	body := dto.Drop{}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/drop/create", bytes.NewReader(bodyBytes))
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

func TestDropCreate_ok(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)

	loggedInUserToken := loginUser(t, ts)

	duration, err := time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	futureTime := currentTime.Add(duration)

	body := dto.Drop{
		CollectionName:     "Col #1",
		Blockchain:         int(models.OtherBlockchain),
		PublicSaleDateTime: futureTime,
		TimeZone:           "America/Adak",
		WebsiteURL:         "http://example.com",
		PublicSalePrice:    1,
		TotalSupply:        5000,
		BlockchainName:     "CustomBlockChain",
		BillingPlan:        int(models.StandardBillingPlan),
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/drop/create", bytes.NewReader(bodyBytes))
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
