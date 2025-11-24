package main

import (
	"log"
	"net/http"
	"os"

	"menu-service/database"
	"menu-service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=menu-db user=postgres password=postgres dbname=menu_db port=5432 sslmode=disable"
	}

	if err := database.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/menu", handlers.CreateMenuItem)
	r.Get("/menu/{id}", handlers.GetMenuItem)
	r.Get("/menu", handlers.GetAllMenuItems)

	log.Printf("Menu-service server starting on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
