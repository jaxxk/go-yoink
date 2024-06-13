package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jaxxk/go-yoink/internal/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var apiCfg *apiConfig

func TestMain(m *testing.M) {
	connStr := os.Getenv("CONN")
	if connStr == "" {
		panic("CONN environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	apiCfg = &apiConfig{
		DB: database.New(db),
	}

	code := m.Run()
	os.Exit(code)
}

func TestHandlerGetUserByApiKey(t *testing.T) {
	// Setup router
	r := chi.NewRouter()
	r.Get("/v1/users", apiCfg.middlewareAuth(apiCfg.HandlerGetUser))

	// Create test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make request with valid API key
	req, err := http.NewRequest("GET", ts.URL+"/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "valid_api_key")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
