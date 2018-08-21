package errors

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
)

const stackSize = 4 << 10 // 4KB

// Handler is an Echo error handler that uses HTTP errors accordingly, and any
// generic error will be interpreted as an internal server error.
func Handler(err error, c echo.Context) {
	log := c.Get("logger").(logger.Logger)

	code := http.StatusInternalServerError
	var msg interface{}
	msg = err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if code == http.StatusInternalServerError {
		stack := make([]byte, stackSize)
		length := runtime.Stack(stack, true)
		log.Err(err).Error("server error", logger.Data{"stack": stack[:length]})
		msg = "internal server error"
	}
	err = c.JSON(code, map[string]interface{}{"error": map[string]interface{}{"message": msg, "status_code": code}})
	if err != nil {
		log.Err(err).Error("error handler json error")
	}
}
