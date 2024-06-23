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
		return
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
		return
	}

	followFeed := database.FollowFeedParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    dbFeed.ID,
	}

	dbFollowFeed, err := cfg.DB.FollowFeed(r.Context(), followFeed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to follow feed")
		return
	}

	type response struct {
		Feed       database.Feed       `json:"feed"`
		FollowFeed database.FeedFollow `json:"follow_feed"`
	}

	// Create a response instance
	resp := response{
		Feed:       dbFeed,
		FollowFeed: dbFollowFeed,
	}

	// Respond with the combined structure
	respondWithJSON(w, http.StatusOK, resp)
}
