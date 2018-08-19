package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	app := &App{}

	e := echo.New()

	e.Use(Middleware(app))

	e.GET("/app", func(c echo.Context) error {
		app, exists := c.Get("app").(*App)
		assert.Equal(t, true, exists, "expected app to exist on context")
		assert.NotNil(t, app, "expected app to be non-nil")
		return c.String(http.StatusOK, "ok")
	})

	req, err := http.NewRequest("GET", "/app", nil)
	require.Nil(t, err, "unexpecetd error when making new request")

	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "incorrect status code")
	assert.Equal(t, "ok", w.Body.String(), "incorrect response")
}
