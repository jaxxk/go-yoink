package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jaxxk/go-yoink/internal/database"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ensure the cancel function is called to release resources

	type parameters struct {
		Name   string `json:"name"`
		ApiKey string `json:"api_key"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to decode")
	}

	var user database.CreateUserParams
	apikey, err := generateApiKey()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to generate api key")
	}

	if params.ApiKey == "" {
		user = database.CreateUserParams{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      params.Name,
			ApiKey:    apikey, // Assuming empty string should be handled by your SQL logic to generate a new ApiKey
		}
	} else {
		user = database.CreateUserParams{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      params.Name,
			ApiKey:    params.ApiKey, // Assuming empty string should be handled by your SQL logic to generate a new ApiKey
		}
	}

	dbUser, err := cfg.DB.CreateUser(ctx, user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create user")
	}

	respondWithJSON(w, http.StatusOK, dbUser)
}

func generateApiKey() (string, error) {
	// Generate a random byte array
	randBytes := make([]byte, 32) // 32 bytes = 256 bits
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		return "", err
	}

	// Create a SHA-256 hash of the random bytes
	hash := sha256.New()
	hash.Write(randBytes)
	hashedBytes := hash.Sum(nil)

	// Encode the hashed bytes to a hexadecimal string
	apiKey := hex.EncodeToString(hashedBytes)
	return apiKey, nil
}
