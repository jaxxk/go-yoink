package main

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jaxxk/go-yoink/internal/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var apiCfg *apiConfig
var ts *httptest.Server

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

	r := chi.NewRouter()

	// Apply middleware
	r.Use(Recoverer)
	r.Use(Logger)

	r.Get("/v1/healthz", handlerReadinessCheck)
	r.Get("/v1/err", handlerErr)
	r.Post("/v1/users", apiCfg.handlerCreateUser)
	r.Get("/v1/users", apiCfg.middlewareAuth(apiCfg.HandlerGetUser))

	ts = httptest.NewServer(r)
	defer ts.Close()

	// Run the tests
	code := m.Run()

	// Clean up
	db.Close()

	os.Exit(code)
}

func TestHandlerGetUserByApiKey(t *testing.T) {
	// Add user
	user := database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "test user",
		ApiKey:    "valid_api_key",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ensure the cancel function is called to release resources
	_, err := apiCfg.DB.CreateUser(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	// Make request with valid API key
	req, err := http.NewRequest("GET", ts.URL+"/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "ApiKey valid_api_key")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
