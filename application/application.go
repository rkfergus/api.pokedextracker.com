package application

import (
	"context"

	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/config"
)

type App struct {
	DB     *pg.DB
	Config *config.Config
}

func New(db *pg.DB, cfg *config.Config) *App {
	return &App{db, cfg}
}

type key struct{}

var contextKey = &key{}

func (app *App) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey, app)
}

func FromContext(ctx context.Context) *App {
	if app, ok := ctx.Value(contextKey).(*App); ok {
		return app
	}
	return &App{}
}
