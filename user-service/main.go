package main

import (
	"log"
	"net/http"
	"os"
	"user-service/database"
	"user-service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=user-db user=postgres password=postgres dbname=user_db port=5432 sslmode=disable"
	}

	if err := database.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", handlers.CreateUser)
	r.Get("/users/{id}", handlers.GetUser)

	log.Printf("User-service server starting on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
