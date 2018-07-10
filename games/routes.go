package games

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	g := r.Group("/games")

	g.GET("", listHandler)
}
