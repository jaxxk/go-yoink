package main

import "net/http"

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	allFeeds, err := cfg.DB.GetAllFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to get all feeds")
		return
	}

	if allFeeds == nil {
		respondWithJSON(w, http.StatusOK, "No feeds")
		return
	}
	respondWithJSON(w, http.StatusOK, allFeeds)
}
