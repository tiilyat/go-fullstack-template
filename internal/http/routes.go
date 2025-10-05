package http

import (
	"github.com/go-chi/chi/v5"
)

func bindRoutes(router chi.Router) {
	router.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler())
	})

	router.HandleFunc("/*", spaHandler())
}
