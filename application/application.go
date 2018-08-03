package application

import (
	"github.com/go-pg/pg/orm"
	"github.com/pokedextracker/api.pokedextracker.com/config"
)

// App contains necessary references that will be persisted throughout the
// application's lifecycle.
type App struct {
	DB     orm.DB
	Config *config.Config
}

// New creates a new instance of App with the given DB and Config.
func New(db orm.DB, cfg *config.Config) *App {
	return &App{db, cfg}
}
