package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthRoute(t *testing.T) {
	r := gin.New()

	RegisterRoutes(r)

	req, err := http.NewRequest("GET", "/health", nil)
	require.Nil(t, err, "unexpecetd error when making new request")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "incorrect status code")
	assert.Equal(t, `{"healthy":true}`, w.Body.String(), "incorrect response")
}
