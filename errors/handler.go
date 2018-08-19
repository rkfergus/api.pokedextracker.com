package errors

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/logger"
)

const stackSize = 4 << 10 // 4KB

func Handler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var msg interface{}
	msg = err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if code == http.StatusInternalServerError {
		log := c.Get("logger").(logger.Logger)
		stack := make([]byte, stackSize)
		length := runtime.Stack(stack, true)
		log.Err(err).Error("server error", logger.Data{"stack": stack[:length]})
		msg = "internal server error"
	}
	c.JSON(code, map[string]interface{}{"error": map[string]interface{}{"message": msg, "status_code": code}})
}
