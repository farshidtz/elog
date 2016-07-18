package logger

import (
	"io"
	"fmt"
	"time"
)

type writer struct {
	io.Writer
	timeFormat string
	prefix     string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(fmt.Sprintf("%s%s", time.Now().Format(w.timeFormat), w.prefix)), b...))
}
