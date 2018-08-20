package users

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/dexes"
	"github.com/pokedextracker/api.pokedextracker.com/errors"
	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func createHandler(c echo.Context) error {
	app := c.Get("app").(*application.App)

	params := createParams{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	// TODO: check for existing username

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), cost)
	if err != nil {
		return err
	}

	xff := c.Request().Header.Get("x-forwarded-for")
	var ip string
	if xff != "" {
		ip = strings.TrimSpace(strings.Split(xff, ",")[0])
	} else {
		ip = c.Request().RemoteAddr
	}

	user := User{
		Username:   params.Username,
		Password:   string(hash),
		FriendCode: params.FriendCode,
		Referrer:   params.Referrer,
		LastIP:     &ip,
	}
	dex := dexes.Dex{
		Title:    params.Title,
		Slug:     params.Title, // TODO: slugify the title
		Shiny:    *params.Shiny,
		GameID:   params.Game,
		Regional: *params.Regional,
	}

	err = app.DB.RunInTransaction(func(tx *pg.Tx) error {
		if _, err := tx.Model(&user).Insert(); err != nil {
			return err
		}
		dex.UserID = user.ID
		_, err := tx.Model(&dex).Insert()
		return err
	})
	if err != nil {
		if e, ok := err.(pg.Error); ok && e.Field('C') == errors.PGUniqueViolationCode {
			return errors.ExistingUsername()
		}
		return err
	}
	err = app.DB.
		Model(&user).
		Relation("Dexes").
		Relation("Dexes.Game").
		Relation("Dexes.Game.GameFamily").
		Where("username = ?", user.Username).
		First()
	if err != nil {
		return err
	}

	// TODO: return session
	return c.JSON(http.StatusOK, user)
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
