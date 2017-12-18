package elog

import (
	"io"
	"time"
)

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

// NewWriter returns an io.Writer that writes timestamps as prefix
func NewWriter(writer io.Writer, timeFormat string) *writer {
	if timeFormat == "" {
		timeFormat = "2006-01-02 15:04:05 "
	}
	return &writer{writer, timeFormat}
}
