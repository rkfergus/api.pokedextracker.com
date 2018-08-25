package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/binder"
	"github.com/stretchr/testify/require"
)

// NewContext returns a new echo.Context, *http.Request, and
// *httptest.ResponseRecorder to be used for tests.
func NewContext(t *testing.T, payload []byte) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	b, err := binder.New()
	require.NoError(t, err)
	e.Binder = b
	req := httptest.NewRequest(echo.GET, "/", bytes.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)
	return c, req, rr
}

// SerializeStruct takes in an interface that is converted into a []byte to be
// used as a payload for requests.
func SerializeStruct(t *testing.T, i interface{}) []byte {
	data, err := json.Marshal(i)
	require.NoError(t, err)
	return data
}
