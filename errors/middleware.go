package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// normal errors
		if err := c.Errors.ByType(gin.ErrorTypePublic).Last(); err != nil {
			status := c.Writer.Status()
			c.JSON(status, gin.H{"error": gin.H{"message": err, "status_code": status}})
		}
		// internal server errors
		if err := c.Errors.ByType(gin.ErrorTypePrivate).Last(); err != nil {
			status := http.StatusInternalServerError
			c.JSON(status, gin.H{"error": gin.H{"message": "internal server error", "status_code": status}})
		}
	}
}
