package config

func loadDevelopmentConfig(cfg *Config) {
	cfg.DatabaseHost = "127.0.0.1"
	cfg.DatabaseName = "pokedex_tracker"
	cfg.DatabaseSSLMode = "disable"
	cfg.DatabaseUser = "pokedex_tracker_user"
	cfg.Environment = "development"
}
