package errors

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	cases := []struct {
		err  error
		code int
		msg  string
	}{
		{ExistingUsername(), http.StatusUnprocessableEntity, "username is already taken"},
		{NotFound("foo"), http.StatusNotFound, "foo not found"},
	}

	for _, tt := range cases {
		err := tt.err.(*echo.HTTPError)
		msg := err.Message.(string)
		assert.Equal(t, tt.code, err.Code)
		assert.Equal(t, tt.msg, msg)
	}
}
