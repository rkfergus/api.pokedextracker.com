package config

import (
	"os"
	"strconv"
)

func loadDevelopmentConfig(cfg *Config) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		cfg.Port = port
	}

	cfg.DatabaseHost = "127.0.0.1"
	cfg.DatabaseName = "pokedex_tracker"
	cfg.DatabaseSSLMode = "disable"
	cfg.DatabaseUser = "pokedex_tracker_user"
	cfg.Environment = "development"
}
