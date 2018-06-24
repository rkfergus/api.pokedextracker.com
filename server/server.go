package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/games"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/pokemon"
	"github.com/rs/zerolog/log"
)

func New(app *application.App) *http.Server {
	r := mux.NewRouter()

	r.Use(logger.Middleware)

	pokemon.RegisterRoutes(r)
	games.RegisterRoutes(r)

	r.NotFoundHandler = logger.Middleware(http.HandlerFunc(notFoundHandler))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: r,
	}

	graceful := make(chan os.Signal, 1)
	signal.Notify(graceful, syscall.SIGINT, syscall.SIGTERM)

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
