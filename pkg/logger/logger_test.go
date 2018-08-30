package logger

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	originalStdout := os.Stdout
	defer func() {
		os.Stdout = originalStdout
	}()
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w

	read := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		read <- buf.String()
	}()

	log := New().
		ID("id123").
		Data(Data{"1": "1", "2": "2"}).
		Root(Data{"4": "4"}).
		Err(errors.New("test error"))

	log.Info("foo", Data{"1": "11", "3": "3"})
	log.Error("foo", Data{"1": "11", "3": "3"})
	log.Warn("foo", Data{"1": "11", "3": "3"})
	log.Debug("foo", Data{"1": "11", "3": "3"})

	err = w.Close()
	require.NoError(t, err)

	line := <-read

	assert.Contains(t, line, `"timestamp":`)
	assert.Contains(t, line, `"hostname":`)
	assert.Contains(t, line, `"id":"id123"`)
	assert.Contains(t, line, `"1":"11"`)
	assert.Contains(t, line, `"2":"2"`)
	assert.Contains(t, line, `"3":"3"`)
	assert.Contains(t, line, `"4":"4"`)
	assert.Contains(t, line, `"error":"test error"`)
	assert.Contains(t, line, `"level":"info"`)
	assert.Contains(t, line, `"message":"foo"`)
}
