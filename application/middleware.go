package application

import (
	"github.com/gin-gonic/gin"
)

func Middleware(app *App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	}
}
