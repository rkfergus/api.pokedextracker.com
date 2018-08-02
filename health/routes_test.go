package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthRoute(t *testing.T) {
	r := gin.New()

	RegisterRoutes(r)

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("unexpecetd error when making new request: %s", err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("incorrect status code\ngot:  %d\nwant: %d", w.Code, http.StatusOK)
	}

	body := w.Body.String()
	wantBody := `{"healthy":true}`
	if body != wantBody {
		t.Errorf("incorrect response\ngot:  %s\nwant: %s", body, wantBody)
	}
}
