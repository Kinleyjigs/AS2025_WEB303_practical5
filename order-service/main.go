package main

import (
	"log"
	"net/http"
	"os"

	"order-service/database"
	"order-service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=order-db user=postgres password=postgres dbname=order_db port=5432 sslmode=disable"
	}

	if err := database.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/orders", handlers.CreateOrder)
	r.Get("/orders/{id}", handlers.GetOrder)
	r.Get("/orders", handlers.GetAllOrders)

	log.Printf("Order-service server starting on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
