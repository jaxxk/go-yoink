package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jaxxk/go-yoink/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CONN := os.Getenv("CONN")
	db, err := sql.Open("postgres", CONN)

	config := apiConfig{
		DB: database.New(db),
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
