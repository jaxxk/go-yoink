package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jaxxk/go-yoink/internal/database"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ensure the cancel function is called to release resources

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to decode")
	}

	user := database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	}

	dbUser, err := cfg.DB.CreateUser(ctx, user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create user")
	}

	respondWithJSON(w, http.StatusOK, dbUser)
}
