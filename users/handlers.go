package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/errors"
)

func createHandler(c *gin.Context) {
	params := createParams{}
	if err := c.Bind(&params); err != nil {
		c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, params.Username)
}

func listHandler(c *gin.Context) {

}

func retrieveHandler(c *gin.Context) {
	app := c.MustGet("app").(*application.App)
	username := c.Param("username")

	user := User{}

	_, err := app.DB.QueryOne(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		if err == pg.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, errors.NotFound("user")).SetType(gin.ErrorTypePublic)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func updateHandler(c *gin.Context) {

}
