package application

import (
	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/config"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/database"
)

// App contains necessary references that will be persisted throughout the
// application's lifecycle.
type App struct {
	DB     *pg.DB
	Config *config.Config
}

// New creates a new instance of App with a DB and Config.
func New() (*App, error) {
	cfg := config.New()
	db, err := database.New(cfg)
	if err != nil {
		return nil, err
	}
	return &App{db, cfg}, nil
}
