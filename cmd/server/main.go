package main

import (
	"log"

	"github.com/TheArchitectEngineer/mirros_api/internal/app"
	"github.com/TheArchitectEngineer/mirros_api/internal/config"
	"github.com/TheArchitectEngineer/mirros_api/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := store.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}

	r := app.NewRouter(db, cfg)
	log.Printf("starting server on %s", cfg.BindAddr)
	if err := r.Run(cfg.BindAddr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}