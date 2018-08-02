package recovery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
)

func TestRecovery(t *testing.T) {
	r := gin.New()
	r.Use(logger.Middleware())
	r.Use(Middleware())

	r.GET("/error", func(c *gin.Context) {
		panic(errors.New("error"))
	})
	r.GET("/string", func(c *gin.Context) {
		panic("string")
	})

	req, err := http.NewRequest("GET", "/error", nil)
	if err != nil {
		t.Fatalf("unexpecetd error when making new request: %s", err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("incorrect recovered status code\ngot:  %d\nwant: %d", w.Code, http.StatusInternalServerError)
	}

	body := w.Body.String()
	wantBody := "internal server error"
	if !strings.Contains(body, wantBody) {
		t.Errorf("incorrect error message\ngot:  %s\nwant: %s", body, wantBody)
	}

	req, err = http.NewRequest("GET", "/string", nil)
	if err != nil {
		t.Fatalf("unexpecetd error when making new request: %s", err)
	}
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected to handle panicking with a non-error")
	}
}
