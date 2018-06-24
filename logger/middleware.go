package logger

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func Middleware(next http.Handler) http.Handler {
	l := New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		rec := &statusRecorder{w, 200}
		id := uuid.NewV4()
		log := l.With().Str("id", id.String()).Logger()
		ctx := log.WithContext(r.Context())
		next.ServeHTTP(rec, r.WithContext(ctx))
		t2 := time.Now()
		log.Info().
			Int("status_code", rec.status).
			Str("method", r.Method).
			Str("path", r.RequestURI).
			Float64("duration", t2.Sub(t1).Seconds()*1000).
			Str("referer", r.Referer()).
			Str("user_agent", r.UserAgent()).
			Msg("request handled")
	})
}
