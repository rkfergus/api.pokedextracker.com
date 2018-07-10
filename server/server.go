package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/games"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/pokedextracker/api.pokedextracker.com/pokemon"
	"github.com/pokedextracker/api.pokedextracker.com/signals"
)

func New(app *application.App) *http.Server {
	log := logger.New()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(logger.Middleware())
	r.Use(gin.Recovery())
	r.Use(application.Middleware(app))

	pokemon.RegisterRoutes(r)
	games.RegisterRoutes(r)

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
	c.JSON(404, gin.H{"error": gin.H{"message": "not found", "status_code": 404}})
}
