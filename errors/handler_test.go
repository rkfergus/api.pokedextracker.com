package errors

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	c, rr := newContext()
	err := errors.New("foo")
	Handler(err, c)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "expected generic errors to be 500s")
	assert.Contains(t, rr.Body.String(), "internal server error", "expected generic errors to have the correct message")

	c, rr = newContext()
	err = echo.NewHTTPError(http.StatusTeapot, "foo")
	Handler(err, c)
	assert.Equal(t, http.StatusTeapot, rr.Code, "expected HTTP errors to be correct")
	assert.Contains(t, rr.Body.String(), "foo", "expected HTTP errors to have the correct message")
}

func newContext() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)
	c.Set("logger", logger.New())
	return c, rr
}
