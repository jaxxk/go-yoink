package main

import (
	"encoding/json"
	"net/http"

	"github.com/jaxxk/go-yoink/internal/database"
)

const DEFAULT_LIMIT int32 = 5

func (cfg *apiConfig) HandlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		AmountOfPosts int `json:"amount_of_posts"`
	}
	var count int32
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Failed to decode")
			return
		}
		count = int32(params.AmountOfPosts)
	} else {
		count = DEFAULT_LIMIT
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  count,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to retrieve posts")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
