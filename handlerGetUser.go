package main

import (
	"net/http"

	"github.com/jaxxk/go-yoink/internal/database"
)

func (cfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
