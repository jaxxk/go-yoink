package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Run the tests
	code := m.Run()
	os.Exit(code)
}

var URL string = "http://localhost:8080"

func TestHandlerCreateUser(t *testing.T) {
	type parameters struct {
		Name   string `json:"name"`
		ApiKey string `json:"api_key"`
	}

	params := parameters{
		Name:   "test user",
		ApiKey: "valid_api_key",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}

	// Make POST request to create user
	req, err := http.NewRequest("POST", URL+"/v1/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandlerAuthUser(t *testing.T) {
	req, err := http.NewRequest("GET", URL+"/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Apikey valid_api_key")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandlerCreateFeed(t *testing.T) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{
		Name: "test feed",
		Url:  "the feed",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", URL+"/v1/feeds", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Apikey valid_api_key")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// assert.Equal(t, http.StatusOK, resp.StatusCode)
}
