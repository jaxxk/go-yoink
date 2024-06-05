package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	r := chi.NewRouter()

	r.Get("/v1/healthz", handlerReadinessCheck)
	r.Get("/v1/err", handlerErr)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
