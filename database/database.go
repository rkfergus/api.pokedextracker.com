package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pokedextracker/api.pokedextracker.com/config"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName, cfg.DatabaseSSLMode)
	conn, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return conn.Set("gorm:auto_preload", true), nil
}
