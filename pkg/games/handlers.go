package games

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/models"
)

type handler struct {
	app *application.App
}

func (h *handler) listHandler(c echo.Context) error {
	var games []models.Game

	err := h.app.DB.
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
