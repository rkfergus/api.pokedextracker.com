package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Data map[string]interface{}

type Logger struct {
	zl   zerolog.Logger
	id   string
	data []Data
	err  error
	root []Data
}

func New() Logger {
	host, _ := os.Hostname()
	release := os.Getenv("RELEASE")

	zl := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("host", host)

	if release != "" {
		zl.Str("release", release)
	}

	return Logger{
		zl:   zl.Logger(),
		data: []Data{},
		root: []Data{},
	}
}

func (log Logger) ID(id string) Logger {
	log.id = id
	return log
}

func (log Logger) Data(data Data) Logger {
	log.data = append(log.data, data)
	return log
}

func (log Logger) Err(err error) Logger {
	log.err = err
	return log
}

func (log Logger) Root(root Data) Logger {
	log.root = append(log.root, root)
	return log
}

func (log Logger) Info(message string, fields ...Data) {
	log.log(log.zl.Info(), message, fields...)
}

func (log Logger) Error(message string, fields ...Data) {
	log.log(log.zl.Error(), message, fields...)
}

func (log Logger) Warn(message string, fields ...Data) {
	log.log(log.zl.Warn(), message, fields...)
}

func (log Logger) Debug(message string, fields ...Data) {
	log.log(log.zl.Debug(), message, fields...)
}

func (log Logger) Fatal(message string, fields ...Data) {
	log.log(log.zl.Fatal(), message, fields...)
}

func (log Logger) log(evt *zerolog.Event, message string, fields ...Data) {
	// Merge data fields
	hasData := false
	data := zerolog.Dict()
	for _, field := range append(log.data, fields...) {
		if len(field) != 0 {
			hasData = true
			data = data.Fields(field)
		}
	}

	// Add root fields
	for _, field := range log.root {
		if len(field) != 0 {
			evt = evt.Fields(field)
		}
	}
	// Add id field
	if log.id != "" {
		evt = evt.Str("id", log.id)
	}
	// Add data field
	if hasData {
		evt = evt.Dict("data", data)
	}
	evt.Int64("nanoseconds", time.Now().UnixNano()).Err(log.err).Msg(message)
}
