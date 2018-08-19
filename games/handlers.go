package games

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/application"
)

func listHandler(c echo.Context) error {
	app := c.Get("app").(*application.App)

	var games []Game

	err := app.DB.
		Model(&games).
		Relation("GameFamily").
		Join("INNER JOIN game_families ON games.game_family_id = game_families.id").
		Where("game_families.published = TRUE").
		OrderExpr("game_families.order DESC, games.order ASC").
		Select()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, games)
}
