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
		log.Fatal().Err(err).Msg("")
	}
	app := application.New(db, cfg)
	srv := server.New(app)

	log.Info().Int("port", cfg.Port).Msg("server started")
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server stopped")
	}
	log.Info().Msg("server stopped")
}
