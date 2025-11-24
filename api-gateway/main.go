package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Route /api/users/* to user-service
	r.Handle("/api/users", proxyTo("http://user-service:8081", "/users"))
	r.Handle("/api/users/*", proxyTo("http://user-service:8081", "/users"))

	// Route /api/menu/* to menu-service
	r.Handle("/api/menu", proxyTo("http://menu-service:8082", "/menu"))
	r.Handle("/api/menu/*", proxyTo("http://menu-service:8082", "/menu"))

	// Route /api/orders/* to order-service
	r.Handle("/api/orders", proxyTo("http://order-service:8083", "/orders"))
	r.Handle("/api/orders/*", proxyTo("http://order-service:8083", "/orders"))

	log.Println("API Gateway starting on :8080")
	log.Println("Routes configured:")
	log.Println("  /api/users   -> user-service:8081")
	log.Println("  /api/menu    -> menu-service:8082")
	log.Println("  /api/orders  -> order-service:8083")

	http.ListenAndServe(":8080", r)
}

func proxyTo(targetURL, pathPrefix string) http.Handler {
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Store the original path for logging
		originalPath := r.URL.Path

		// Remove /api prefix and set the new path
		newPath := strings.TrimPrefix(r.URL.Path, "/api")
		r.URL.Path = newPath

		// Log the routing
		log.Printf("Routing %s %s -> %s%s", r.Method, originalPath, targetURL, newPath)

		// Serve the request
		proxy.ServeHTTP(w, r)
	})
}
