package users

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/errors"
)

func createHandler(c echo.Context) error {
	params := createParams{}
	if err := c.Bind(&params); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, params.Username)
}

func listHandler(c echo.Context) error {
	return nil
}

func retrieveHandler(c echo.Context) error {
	app := c.Get("app").(*application.App)
	username := c.Param("username")

	user := User{}

	err := app.DB.
		Model(&user).
		Relation("Dexes", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("dexes.date_created ASC"), nil
		}).
		Relation("Dexes.Game").
		Relation("Dexes.Game.GameFamily").
		Where("username = ?", username).
		First()
	if err != nil {
		if err == pg.ErrNoRows {
			return errors.NotFound("user")
		}
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func updateHandler(c echo.Context) error {
	return nil
}
