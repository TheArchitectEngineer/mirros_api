package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	BindAddr    string
	// Useful defaults for tests/dev
	ShutdownTimeout time.Duration
}

func Load() (*Config, error) {
	// Load .env if present (no error if missing)
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL:     getenv("DATABASE_URL", "postgres://postgres:password@localhost:5432/mirros_dev?sslmode=disable"),
		JWTSecret:       getenv("JWT_SECRET", "replace-me-in-prod"),
		BindAddr:        getenv("BIND_ADDR", ":8080"),
		ShutdownTimeout: 10 * time.Second,
	}
	return cfg, nil
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}