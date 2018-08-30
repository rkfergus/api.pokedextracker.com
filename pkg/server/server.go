package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	elog "github.com/labstack/gommon/log"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/binder"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/errors"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/games"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/health"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/logger"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/pokemon"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/recovery"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/signals"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/users"
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

	health.RegisterRoutes(e)
	games.RegisterRoutes(e, app)
	pokemon.RegisterRoutes(e, app)
	users.RegisterRoutes(e, app)

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
