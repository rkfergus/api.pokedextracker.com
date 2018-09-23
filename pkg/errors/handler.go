package errors

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/logger"
)

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
		log.Err(err).Error("server error")
		msg = "internal server error"
	}
	err = c.JSON(code, map[string]interface{}{"error": map[string]interface{}{"message": msg, "status_code": code}})
	if err != nil {
		log.Err(err).Error("error handler json error")
	}
}
