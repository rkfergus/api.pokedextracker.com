package users

import "github.com/labstack/echo"

// RegisterRoutes takes in an Echo router and resgiters routes onto it.
func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/users")

	g.POST("", createHandler)
	g.GET("", listHandler)
	g.GET("/:username", retrieveHandler)
	g.POST("/:id", updateHandler)
}
