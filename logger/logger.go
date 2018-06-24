package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.CallerFieldName = "file"
}

func New() zerolog.Logger {
	host, _ := os.Hostname()
	log := zerolog.New(os.Stdout)
	return log.With().
		Timestamp().
		Str("host", host).
		Logger()
}
