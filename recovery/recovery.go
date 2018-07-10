package recovery

import (
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
				err, ok := r.(error)
				if ok {
					log = log.Err(err)
				} else {
					log = log.Data(map[string]interface{}{"error": r})
				}
				log.Error("server error", map[string]interface{}{"stack": debug.Stack()})
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": gin.H{"message": "internal server error", "status_code": http.StatusInternalServerError}})
			}
		}()
		c.Next()
	}
}
