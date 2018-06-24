package config

import (
	"os"
)

type Config struct {
	DatabaseHost     string
	DatabaseName     string
	DatabasePassword string
	DatabasePort     int
	DatabaseSSLMode  string
	DatabaseUser     string
	Environment      string
	Port             int
}

func New() *Config {
	cfg := &Config{
		Port:         8648,
		DatabasePort: 5432,
	}

	switch os.Getenv("APP_ENV") {
	case "development", "":
		loadDevelopmentConfig(cfg)
	case "staging":
		loadStagingConfig(cfg)
	case "production":
		loadProductionConfig(cfg)
	}

	return cfg
}
