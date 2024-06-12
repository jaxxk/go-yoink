package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// Logger is a middleware that logs the start and end of each request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
	})
}

// Recoverer is a middleware that recovers from panics and writes a 500 error response.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Middleware authentication
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyString := r.Header.Get("Authorization")
		if apiKeyString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse apiKeyString
		apiKey := parseApiKey(apiKeyString)

		// Fetch the user from the database using the API key
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler with the user
		handler(w, r, user)
	}
}

func parseApiKey(s string) string {
	strList := strings.Split(s, " ")
	return strList[1]
}
