package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/pokemon"
	"github.com/pokedextracker/api.pokedextracker.com/signals"
	"github.com/rs/zerolog/log"
)

func New(app *application.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(logger.Middleware())
	r.Use(gin.Recovery())
	r.Use(application.Middleware(app))

	pokemon.RegisterRoutes(r)
	// games.RegisterRoutes(r)

	// r.NotFoundHandler = logger.Middleware(http.HandlerFunc(notFoundHandler))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: r,
	}

	graceful := signals.Setup()

	go func() {
		<-graceful
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("server shutdown")
		}
	}()

	return srv
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintln(w, `{"error":{"message":"not found","status_code":404}}`)
}
