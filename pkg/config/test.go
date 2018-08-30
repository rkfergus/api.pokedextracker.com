package config

func loadTestConfig(cfg *Config) {
	cfg.DatabaseHost = "127.0.0.1"
	cfg.DatabaseName = "pokedex_tracker_test"
	cfg.DatabaseSSLMode = "disable"
	cfg.DatabaseUser = "pokedex_tracker_admin"
	cfg.Environment = "test"
}
