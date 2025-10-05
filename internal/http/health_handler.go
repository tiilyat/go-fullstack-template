package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthHandler returns a handler for health check endpoint.
func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{
			"message": "ok",
		}); err != nil {
			log.Printf("Failed to encode health response: %v", err)
		}
	}
}
