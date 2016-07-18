package logger

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	*log.Logger
	debugger *log.Logger
}

func (l *Logger) Errorf(format string, a ...interface{}) error {
	if l.debugger != nil {
		l.debugger.Output(2, fmt.Sprintf(format, a...))
	}
	return fmt.Errorf(format, a...)
}

func (l *Logger) Debug(a ...interface{}) {
	if l.debugger != nil {
		l.debugger.Output(2, fmt.Sprint(a...))
	}
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.debugger != nil {
		l.debugger.Output(2, fmt.Sprintf(format, a...))
	}
}

func (l *Logger) Debugln(a ...interface{}) {
	if l.debugger != nil {
		l.debugger.Output(2, fmt.Sprintln(a...))
	}
}


func NewLogger(out io.Writer, timeFormat, prefix string, debug bool) Logger {
	var logger Logger
	timeFormat = fmt.Sprintf("%s ", timeFormat)

	if debug {
		logger.Logger = log.New(&writer{out, timeFormat, fmt.Sprintf("[%s] ", prefix)}, "", log.Lshortfile)
		logger.debugger = log.New(&writer{out, timeFormat, fmt.Sprintf("[%s-debug] ", prefix)}, "", log.Lshortfile)
	} else {
		logger.Logger = log.New(&writer{out, timeFormat, fmt.Sprintf("[%s] ", prefix)}, "", 0)
	}
	return logger
}
