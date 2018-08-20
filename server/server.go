package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	elog "github.com/labstack/gommon/log"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/binder"
	"github.com/pokedextracker/api.pokedextracker.com/errors"
	"github.com/pokedextracker/api.pokedextracker.com/games"
	"github.com/pokedextracker/api.pokedextracker.com/health"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/pokemon"
	"github.com/pokedextracker/api.pokedextracker.com/recovery"
	"github.com/pokedextracker/api.pokedextracker.com/signals"
	"github.com/pokedextracker/api.pokedextracker.com/users"
)

// New returns a new HTTP server with the registered routes.
func New(app *application.App) (*http.Server, error) {
	log := logger.New()
	e := echo.New()

	e.Logger.SetLevel(elog.OFF)

	b, err := binder.New()
	if err != nil {
		return nil, err
	}
	e.Binder = b

	e.Use(logger.Middleware())
	e.Use(recovery.Middleware())
	e.Use(application.Middleware(app))

	games.RegisterRoutes(e)
	health.RegisterRoutes(e)
	pokemon.RegisterRoutes(e)
	users.RegisterRoutes(e)

	echo.NotFoundHandler = notFoundHandler
	e.HTTPErrorHandler = errors.Handler

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: e,
	}

	graceful := signals.Setup()

	go func() {
		<-graceful
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Err(err).Error("server shutdown")
		}
	}()

	return srv, nil
}

func notFoundHandler(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotFound, "not found")
}
