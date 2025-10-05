package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewServer creates and configures the HTTP router with middleware.
func NewServer(config ServeConfig) chi.Router {
	router := chi.NewRouter()

	// Request timeout middleware
	router.Use(middleware.Timeout(60 * time.Second))

	// Logger middleware
	router.Use(middleware.Logger)

	// Recoverer middleware for panic handling
	router.Use(middleware.Recoverer)

	bindRoutes(router)

	return router
}
