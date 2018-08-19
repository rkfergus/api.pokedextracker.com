package recovery

import (
	"fmt"

	"github.com/labstack/echo"
)

func Middleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err := fmt.Errorf("%v", r)
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
