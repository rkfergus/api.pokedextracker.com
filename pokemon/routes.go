package pokemon

import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/pokemon")

	g.GET("", listHandler)
	g.GET("/:id", retrieveHandler)
}
