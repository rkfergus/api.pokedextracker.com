package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func NotFound(resource string) error {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}
