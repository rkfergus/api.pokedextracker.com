package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Middleware() gin.HandlerFunc {
	l := New()
	return func(c *gin.Context) {
		t1 := time.Now()
		id := uuid.NewV4()
		log := l.ID(id.String())
		c.Set("logger", log)
		c.Next()
		t2 := time.Now()
		log.Root(Data{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"duration":    t2.Sub(t1).Seconds() * 1000,
			"referer":     c.Request.Referer(),
			"user_agent":  c.Request.UserAgent(),
		}).Info("request handled")
	}
}
