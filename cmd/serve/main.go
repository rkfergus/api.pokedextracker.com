package main

import (
	"net/http"

	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/server"
)

func main() {
	log := logger.New()

	app, err := application.New()
	if err != nil {
		log.Err(err).Fatal("application error")
	}
	srv, err := server.New(app)
	if err != nil {
		log.Err(err).Fatal("server error")
	}

	log.Info("server started", logger.Data{"port": app.Config.Port})
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("server stopped")
	}
	log.Info("server stopped")
}
