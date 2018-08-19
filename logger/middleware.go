package logger

import (
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func Middleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	l := New()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t1 := time.Now()
			id := uuid.NewV4()
			log := l.ID(id.String())
			c.Set("logger", log)
			if err := next(c); err != nil {
				c.Error(err)
			}
			t2 := time.Now()
			log.Root(Data{
				"status_code": c.Response().Status,
				"method":      c.Request().Method,
				"path":        c.Request().URL.Path,
				"duration":    t2.Sub(t1).Seconds() * 1000,
				"referer":     c.Request().Referer(),
				"user_agent":  c.Request().UserAgent(),
			}).Info("request handled")
			return nil
		}
	}
}
