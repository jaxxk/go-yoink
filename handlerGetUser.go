package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

func (cfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKeyString := r.Header.Get("authorization")
	if apiKeyString == "" {
		respondWithError(w, http.StatusBadRequest, "missing api key")
	}
	apiKey := parseApiKey(apiKeyString)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ensure the cancel function is called to release resources

	dbUser, err := cfg.DB.GetUserByApiKey(ctx, apiKey)
	if err != nil {
		log.Printf("API key is: %s", apiKey)
		respondWithError(w, http.StatusBadRequest, "failed to retireve user from api key")
	}

	respondWithJSON(w, http.StatusOK, dbUser)
}

func parseApiKey(s string) string {
	strList := strings.Split(s, " ")
	return strList[1]
}
