package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	serverMux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: serverMux,
	}
	server.ListenAndServe()
}
