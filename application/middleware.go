package application

import (
	"github.com/labstack/echo"
)

// Middleware returns a new Echo middleware function that attaches the given App
// instance to the request context.
func Middleware(app *App) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", app)
			return next(c)
		}
	}
}
