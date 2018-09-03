package users

import (
	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
)

// RegisterRoutes takes in an Echo router and registers routes onto it.
func RegisterRoutes(e *echo.Echo, app *application.App) {
	g := e.Group("/users")

	h := handler{app}

	g.POST("", h.create)
	g.GET("", h.list)
	g.GET("/:username", h.retrieve)
	g.POST("/:id", h.update)
}
