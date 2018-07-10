package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/application"
)

func listHandler(c *gin.Context) {
	app := c.MustGet("app").(*application.App)

	var games []Game

	app.DB.
		Joins("INNER JOIN game_families ON games.game_family_id = game_families.id").
		Where("game_families.published = TRUE").
		Order("game_families.order DESC, games.order ASC").
		Find(&games)

	c.JSON(http.StatusOK, games)
}
