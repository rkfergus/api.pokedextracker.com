package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/config"
)

func New(cfg *config.Config) (*pg.DB, error) {
	addr := fmt.Sprintf("%s:%d", cfg.DatabaseHost, cfg.DatabasePort)
	conn := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		Database: cfg.DatabaseName,
	})
	// Ensure the database can connect
	var n int
	_, err := conn.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
