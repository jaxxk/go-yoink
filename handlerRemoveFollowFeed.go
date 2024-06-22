package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jaxxk/go-yoink/internal/database"
)

func (cfg *apiConfig) HandlerRemoveFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := chi.URLParam(r, "feedFollowID")
	log.Println(feedFollowID)
	err := cfg.DB.Unfollow(r.Context(), feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to unfollow")
		return
	}
	respondWithJSON(w, http.StatusOK, "unfollowed")
}
