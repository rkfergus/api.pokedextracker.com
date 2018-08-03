package recovery

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log := c.MustGet("logger").(logger.Logger)
				err := fmt.Errorf("%v", r)
				log.Err(err).Error("server error", logger.Data{"stack": debug.Stack()})
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": gin.H{"message": "internal server error", "status_code": http.StatusInternalServerError}})
			}
		}()
		c.Next()
	}
}
