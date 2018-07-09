package pokemon

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	g := r.Group("/pokemon")

	g.GET("", listHandler)
	// g.GET("/:id", retrieveHandler)
}
