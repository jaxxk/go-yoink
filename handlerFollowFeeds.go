package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jaxxk/go-yoink/internal/database"
)

func (cfg *apiConfig) HandlerFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to decode")
		return
	}

	followFeed := database.FollowFeedParams{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	}

	dbFollowFeed, err := cfg.DB.FollowFeed(r.Context(), followFeed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to follow feed")
	}

	respondWithJSON(w, http.StatusOK, dbFollowFeed)
}
func (cfg *apiConfig) HandlerGetFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	listOfFollowFeed, err := cfg.DB.GetAllFollowFeedsForUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to retrieve followed feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, listOfFollowFeed)
}
