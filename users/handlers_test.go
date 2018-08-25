package users

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/errors"
	"github.com/pokedextracker/api.pokedextracker.com/internal/test"
	"github.com/pokedextracker/api.pokedextracker.com/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateHandler(t *testing.T) {
	h := newHandler(t)

	t.Run("saves rows on successful creation", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		var params createParams
		err := createParamsFactory.Construct(&params)
		require.NoError(tt, err)

		payload := test.SerializeStruct(tt, params)
		c, _, rr := test.NewContext(tt, payload)

		err = h.createHandler(c)
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
		var params createParams
		err := createParamsFactory.Construct(&params)
		require.NoError(tt, err)
		params.Username = ""

		payload := test.SerializeStruct(tt, params)
		c, _, _ := test.NewContext(tt, payload)

		err = h.createHandler(c)
		assert.Contains(tt, err.Error(), "required")
	})

	t.Run("returns error on when it's an existing username", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		var params createParams
		err := createParamsFactory.Construct(&params)
		require.NoError(tt, err)

		payload := test.SerializeStruct(tt, params)
		c, _, _ := test.NewContext(tt, payload)

		err = h.createHandler(c)
		require.NoError(tt, err)

		payload = test.SerializeStruct(tt, params)
		c, _, _ = test.NewContext(tt, payload)

		err = h.createHandler(c)
		assert.Equal(tt, err, errors.ExistingUsername())
	})

	t.Run("save IP address from X-Forwarded-For header", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)
		xff := "123.123.123.123"
		var params createParams
		err := createParamsFactory.Construct(&params)
		require.NoError(tt, err)

		payload := test.SerializeStruct(tt, params)
		c, req, _ := test.NewContext(tt, payload)
		req.Header.Add("X-Forwarded-For", xff)

		err = h.createHandler(c)
		assert.NoError(tt, err)

		var user models.User
		err = h.app.DB.Model(&user).First()
		require.NoError(tt, err)

		assert.Equal(tt, xff, *user.LastIP)
	})
}

func newHandler(t *testing.T) handler {
	t.Helper()

	app, err := application.New()
	require.NoError(t, err)
	return handler{app}
}
