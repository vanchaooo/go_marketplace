package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr string
	DatabaseURL string
}

func Load() (Config, error) {
	cfg := Config {
		HTTPAddr: os.Getenv("HTTP_ADDR"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.HTTPAddr == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR is not set")
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is not set")
	}
	return cfg, nil
}