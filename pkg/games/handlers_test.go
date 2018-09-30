package games

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pokedextracker/api.pokedextracker.com/internal/test"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListHandler(t *testing.T) {
	h := newHandler(t)

	t.Run("lists published games order correctly", func(tt *testing.T) {
		test.TruncateTables(tt, h.app.DB)

		c, rr := test.NewContext(tt, nil)

		err := h.list(c)
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, rr.Code)

		var response []models.Game
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		require.NoError(tt, err)
		assert.Len(tt, response, 8)
		assert.Equal(tt, models.UltraSunUltraMoonID, response[0].GameFamily.ID)
		assert.Equal(tt, models.XYID, response[len(response)-1].GameFamily.ID)
	})
}

func newHandler(t *testing.T) handler {
	t.Helper()

	app, err := application.New()
	require.NoError(t, err)
	return handler{app}
}
