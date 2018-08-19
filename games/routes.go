package games

import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/games")

	g.GET("", listHandler)
}
