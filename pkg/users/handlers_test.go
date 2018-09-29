package users

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pokedextracker/api.pokedextracker.com/internal/factories"
	"github.com/pokedextracker/api.pokedextracker.com/internal/test"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/errors"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	firstUser  = factories.User.MustCreate().(*models.User)
	secondUser = factories.User.MustCreate().(*models.User)
)

func TestCreateHandler(t *testing.T) {
	h := newHandler(t)

	t.Run("saves rows on successful creation", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		params := createParamsFactory.MustCreate().(*createParams)

		payload := test.SerializeStruct(tt, params)
		c, rr := test.NewContext(tt, payload)

		err := h.create(c)
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, rr.Code)

		var user models.User
		var dex models.Dex
		err = h.app.DB.Model(&user).First()
		require.NoError(tt, err)
		err = h.app.DB.Model(&dex).First()
		require.NoError(tt, err)

		var response models.User
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		require.NoError(tt, err)
		assert.Equal(tt, response.ID, user.ID)
		assert.Equal(tt, response.Dexes[0].ID, dex.ID)
	})

	t.Run("returns error on validation error", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		params := createParamsFactory.MustCreate().(*createParams)
		params.Username = ""

		payload := test.SerializeStruct(tt, params)
		c, _ := test.NewContext(tt, payload)

		err := h.create(c)
		assert.Contains(tt, err.Error(), "required")
	})

	t.Run("returns error on when it's an existing username", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		params := createParamsFactory.MustCreate().(*createParams)

		payload := test.SerializeStruct(tt, params)
		c, _ := test.NewContext(tt, payload)

		err := h.create(c)
		require.NoError(tt, err)

		payload = test.SerializeStruct(tt, params)
		c, _ = test.NewContext(tt, payload)

		err = h.create(c)
		assert.Equal(tt, err, errors.ExistingUsername())
	})

	t.Run("save IP address from X-Forwarded-For header", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		xff := "123.123.123.123"
		params := createParamsFactory.MustCreate().(*createParams)

		payload := test.SerializeStruct(tt, params)
		c, _ := test.NewContext(tt, payload)
		c.Request().Header.Add("X-Forwarded-For", xff)

		err := h.create(c)
		assert.NoError(tt, err)

		var user models.User
		err = h.app.DB.Model(&user).First()
		require.NoError(tt, err)

		assert.Equal(tt, xff, *user.LastIP)
	})
}

func TestListHandler(t *testing.T) {
	h := newHandler(t)

	t.Run("lists users on success", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		users := []*models.User{firstUser, secondUser}
		err := h.app.DB.Insert(&users)
		require.NoError(tt, err)

		c, rr := test.NewContext(tt, nil)
		c.QueryParams().Set("limit", "1")

		err = h.list(c)
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, rr.Code)

		var response []models.User
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		require.NoError(tt, err)
		assert.Len(tt, response, 1)
	})

	t.Run("returns error on validation error", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)

		c, _ := test.NewContext(tt, nil)
		c.QueryParams().Set("limit", "101")

		err := h.list(c)
		assert.Contains(tt, err.Error(), "must be less")
	})

	t.Run("returns empty array when no users are matched", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		users := []*models.User{firstUser, secondUser}
		err := h.app.DB.Insert(&users)
		require.NoError(tt, err)

		c, rr := test.NewContext(tt, nil)
		c.QueryParams().Set("offset", "2")

		err = h.list(c)
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, rr.Code)

		var response []models.User
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		require.NoError(tt, err)
		assert.Len(tt, response, 0)
	})
}

func TestRetrieveHandler(t *testing.T) {
	h := newHandler(t)

	t.Run("retrieves user on success", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		err := h.app.DB.Insert(firstUser)
		require.NoError(tt, err)

		c, rr := test.NewContext(tt, nil)
		c.SetParamNames("username")
		c.SetParamValues(firstUser.Username)

		err = h.retrieve(c)
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, rr.Code)

		var response models.User
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		require.NoError(tt, err)
		assert.Equal(tt, response.ID, firstUser.ID)
	})

	t.Run("returns 404 if user isn't found", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)

		c, _ := test.NewContext(tt, nil)
		c.SetParamNames("username")
		c.SetParamValues("foo")

		err := h.retrieve(c)
		assert.Equal(tt, err, errors.NotFound("user"))
	})
}

func newHandler(t *testing.T) handler {
	t.Helper()

	app, err := application.New()
	require.NoError(t, err)
	return handler{app}
}
