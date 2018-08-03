package recovery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecovery(t *testing.T) {
	r := gin.New()
	r.Use(logger.Middleware())
	r.Use(Middleware())

	r.GET("/error", func(c *gin.Context) { panic(errors.New("error")) })
	r.GET("/string", func(c *gin.Context) { panic("string") })
	r.GET("/int", func(c *gin.Context) { panic(1) })

	paths := []string{"/error", "/string", "/int"}

	for _, path := range paths {
		req, err := http.NewRequest("GET", path, nil)
		require.Nil(t, err, "unexpecetd error when making new request")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "incorrect recovered status code")
		assert.Contains(t, w.Body.String(), "internal server error", "incorrect error message")
	}
}
