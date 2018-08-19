package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/config"
)

// New initializes a new database struct.
func New(cfg *config.Config) (*pg.DB, error) {
	addr := fmt.Sprintf("%s:%d", cfg.DatabaseHost, cfg.DatabasePort)
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		Database: cfg.DatabaseName,
	})
	// Ensure the database can connect
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		fmt.Println("======")
		fmt.Printf("%s %s\n", time.Since(event.StartTime), query)
		fmt.Println("======")
	})
	return db, nil
}
