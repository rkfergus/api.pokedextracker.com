package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Constants for Postgres error codes as found here:
// https://www.postgresql.org/docs/10/static/errcodes-appendix.html.
const (
	PGUniqueViolationCode = "23505"
)

func ExistingUsername() error {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, "username is already taken")
}

func NotFound(resource string) error {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}
