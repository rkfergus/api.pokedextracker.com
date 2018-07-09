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
		log := l.With().Str("id", id.String()).Logger()
		c.Set("logger", log)
		c.Next()
		t2 := time.Now()
		log.Info().
			Int("status_code", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Float64("duration", t2.Sub(t1).Seconds()*1000).
			Str("referer", c.Request.Referer()).
			Str("user_agent", c.Request.UserAgent()).
			Msg("request handled")
	}
}
