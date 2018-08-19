package main

import (
	"net/http"

	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/config"
	"github.com/pokedextracker/api.pokedextracker.com/database"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/server"
)

func main() {
	log := logger.New()

	cfg := config.New()
	db, err := database.New(cfg)
	if err != nil {
		log.Err(err).Fatal("database error")
	}
	app := application.New(db, cfg)
	srv := server.New(app)

	log.Info("server started", logger.Data{"port": cfg.Port})
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("server stopped")
	}
	log.Info("server stopped")
}
