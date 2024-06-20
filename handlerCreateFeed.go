package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jaxxk/go-yoink/internal/database"
)

func (cfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot decode params")
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}

	dbFeed, err := cfg.DB.CreateFeed(r.Context(), feed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot create feed")
	}

	respondWithJSON(w, http.StatusOK, dbFeed)

}
