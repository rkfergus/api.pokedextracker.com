package errors

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/logger"
)

const stackSize = 4 << 10 // 4KB

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Handler is an Echo error handler that uses HTTP errors accordingly, and any
// generic error will be interpreted as an internal server error.
func Handler(err error, c echo.Context) {
	log := logger.FromContext(c)

	code := http.StatusInternalServerError
	var msg interface{}
	msg = err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if code == http.StatusInternalServerError {
		stack := make([]byte, stackSize)
		if err, ok := err.(stackTracer); ok {
			st := err.StackTrace()
			stack = []byte(fmt.Sprintf("%+v", st))
		} else {
			_ = runtime.Stack(stack, true)
		}
		log.Err(err).Error("server error", logger.Data{"stack": stack})
		msg = "internal server error"
	}
	err = c.JSON(code, map[string]interface{}{"error": map[string]interface{}{"message": msg, "status_code": code}})
	if err != nil {
		log.Err(err).Error("error handler json error")
	}
}
