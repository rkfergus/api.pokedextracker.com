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

// ExistingUsername returns a 422 error when a username has already been taken.
func ExistingUsername() error {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, "username is already taken")
}

// NotFound returns a 404 error with a message indicating the given resource.
func NotFound(resource string) error {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}
