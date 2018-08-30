package pokemon

import (
	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
)

// RegisterRoutes takes in an Echo router and registers routes onto it.
func RegisterRoutes(e *echo.Echo, app *application.App) {
	g := e.Group("/pokemon")

	h := handler{app}

	g.GET("", h.listHandler)
	g.GET("/:id", h.retrieveHandler)
}
