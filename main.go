package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jaxxk/go-yoink/internal/database"
	"github.com/jaxxk/go-yoink/worker"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

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

	if err != nil {
		log.Fatal("Failed opening postgres")
	}

	config := apiConfig{
		DB: database.New(db),
	}

	PORT := os.Getenv("PORT")
	r := chi.NewRouter()

	// apply middleware
	r.Use(Recoverer)
	r.Use(Logger)

	r.Get("/v1/healthz", handlerReadinessCheck)
	r.Get("/v1/err", handlerErr)
	r.Post("/v1/users", config.handlerCreateUser)
	r.Get("/v1/users", config.middlewareAuth(config.HandlerGetUser))
	r.Post("/v1/feeds", config.middlewareAuth(config.HandlerCreateFeed))
	r.Post("/v1/feed_follows", config.middlewareAuth(config.HandlerFollowFeed))
	r.Get("/v1/feed_follows", config.middlewareAuth(config.HandlerGetFollowFeed))
	r.Delete("/v1/feed_follows/{feedFollowID}", config.middlewareAuth(config.HandlerRemoveFollowFeed))
	r.Get("/v1/feeds", config.handlerGetFeeds)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}

	go worker.StartWorking(config.DB, worker.FetchRSSFeed)
	log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
