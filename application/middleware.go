package application

import (
	"github.com/gin-gonic/gin"
)

// Middleware returns a new gin middleware function that attaches the given App
// instance to the request context.
func Middleware(app *App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	}
}
