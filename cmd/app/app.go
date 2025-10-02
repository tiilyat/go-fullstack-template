package main

import (
	"log"

	"github.com/tiilyat/embed-go-front/internal/config"
	"github.com/tiilyat/embed-go-front/internal/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	if err := http.Serve(http.ServeConfig{
		Env:               cfg.Env,
		Port:              cfg.Port,
		UIDevMode:         cfg.UIDevMode,
		UIDevServerURL:    cfg.UIDevServerURL,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
