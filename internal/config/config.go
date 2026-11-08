package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr string
	DatabaseURL string
	JWTSecret string
}

func Load() (Config, error) {
	cfg := Config {
		HTTPAddr: os.Getenv("HTTP_ADDR"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.HTTPAddr == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR is not set")
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is not set")
	}
	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("JWT is not set")
	}
	return cfg, nil
}