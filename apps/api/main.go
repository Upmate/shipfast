package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/Upmate/shipfast/apps/api/internal/handler"
	"github.com/Upmate/shipfast/apps/api/internal/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/health", handler.Health)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/products", handler.ListProducts)
		r.Get("/products/{id}", handler.GetProduct)
	})

	log.Println("API server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
