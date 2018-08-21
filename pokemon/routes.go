package pokemon

import "github.com/labstack/echo"

// RegisterRoutes takes in an Echo router and registers routes onto it.
func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/pokemon")

	g.GET("", listHandler)
	g.GET("/:id", retrieveHandler)
}
