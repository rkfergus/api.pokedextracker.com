package binder

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type params struct {
	Hello string `json:"hello" mod:"trim" validate:"token"`
}

var (
	json    = `{"hello":" world "}`
	errJSON = `{"hello":"-"}`
)

func TestNew(t *testing.T) {
	b := New()
	assert.NotNil(t, b)

	c := newContext(json, "invalid")
	p := params{}
	err := b.Bind(&p, c)
	assert.Contains(t, err.Error(), "Unsupported Media Type", "expected Bind to use the default binder")

	c = newContext(json, echo.MIMEApplicationJSON)
	p = params{}
	err = b.Bind(&p, c)
	assert.NoError(t, err)
	assert.Equal(t, "world", p.Hello, "expected Bind to modify the payload")

	c = newContext(errJSON, echo.MIMEApplicationJSON)
	p = params{}
	err = b.Bind(&p, c)
	assert.Contains(t, err.Error(), "alpha-numeric and underscore", "expected Bind to validate the payload")
}

func newContext(payload, mime string) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, mime)
	rr := httptest.NewRecorder()
	return e.NewContext(req, rr)
}
