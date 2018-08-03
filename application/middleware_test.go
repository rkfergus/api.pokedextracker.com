package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	app := &App{}

	r := gin.New()

	r.Use(Middleware(app))

	r.GET("/app", func(c *gin.Context) {
		app, exists := c.Get("app")
		assert.Equal(t, true, exists, "expected app to exist on context")
		assert.NotNil(t, app, "expected app to be non-nil")
		c.String(http.StatusOK, "ok")
	})

	req, err := http.NewRequest("GET", "/app", nil)
	require.Nil(t, err, "unexpecetd error when making new request")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "incorrect status code")
	assert.Equal(t, "ok", w.Body.String(), "incorrect response")
}
