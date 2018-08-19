package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/errors"
	"github.com/pokedextracker/api.pokedextracker.com/games"
	"github.com/pokedextracker/api.pokedextracker.com/health"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/pokemon"
	"github.com/pokedextracker/api.pokedextracker.com/recovery"
	"github.com/pokedextracker/api.pokedextracker.com/signals"
	"github.com/pokedextracker/api.pokedextracker.com/users"

	// load our custom validators
	_ "github.com/pokedextracker/api.pokedextracker.com/validators"
)

// New returns a new HTTP server with the registered routes.
func New(app *application.App) *http.Server {
	log := logger.New()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(logger.Middleware())
	r.Use(recovery.Middleware())
	r.Use(application.Middleware(app))
	r.Use(errors.Middleware())

	games.RegisterRoutes(r)
	health.RegisterRoutes(r)
	pokemon.RegisterRoutes(r)
	users.RegisterRoutes(r)

	r.NoRoute(notFoundHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: r,
	}

	graceful := signals.Setup()

	go func() {
		<-graceful
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Err(err).Error("server shutdown")
		}
	}()

	return srv
}

func notFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "not found", "status_code": http.StatusNotFound}})
}
