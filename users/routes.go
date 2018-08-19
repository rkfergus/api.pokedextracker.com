package users

import "github.com/gin-gonic/gin"

// RegisterRoutes takes in a Gin router and resgiters routes onto it.
func RegisterRoutes(r *gin.Engine) {
	g := r.Group("/users")

	g.POST("", createHandler)
	g.GET("", listHandler)
	g.GET("/:username", retrieveHandler)
	g.POST("/:id", updateHandler)
}
