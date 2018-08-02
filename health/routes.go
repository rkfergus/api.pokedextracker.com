package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", healthHandler)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"healthy": true})
}
