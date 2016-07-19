package elog

import (
	"fmt"
	"log"
	"os"
)

// Logger
type Logger struct {
	log   *log.Logger
	debug *log.Logger
}

// New creates a new Logger.
// prefix is is the prefix used when logging with Print, Fatal, and Panic methods
// config is the configuration struct, with optional parameters.
// if config is set to nil, the default configuration will be used:
//	Writer = os.Stdout
//	TimeFormat = "2006/01/02 15:04:05"
//	DebugEnvVar = "DEBUG"
//	DebugPrefix = "[debug] "
//	DebugTrace = ShortFile
func New(prefix string, config *Config) *Logger {
	conf := initConfig(config)

	var logger Logger
	if *conf.DebugEnabled {
		logger.log = log.New(&writer{conf.Writer, conf.TimeFormat}, prefix, conf.DebugTrace)
		logger.debug = log.New(&writer{conf.Writer, conf.TimeFormat}, conf.DebugPrefix, conf.DebugTrace)
	} else {
		logger.log = log.New(&writer{conf.Writer, conf.TimeFormat}, prefix, 0)
	}
	return &logger
}

// Print prints to the logger with arguments similar to fmt.Print
func (l *Logger) Print(a ...interface{}) {
	l.log.Output(2, fmt.Sprint(a...))
}

// Printf prints to the logger with arguments similar to fmt.Printf
func (l *Logger) Printf(format string, a ...interface{}) {
	l.log.Output(2, fmt.Sprintf(format, a...))
}

// Println prints to the logger with arguments similar to fmt.Println
func (l *Logger) Println(a ...interface{}) {
	l.log.Output(2, fmt.Sprintln(a...))
}

// Fatal is a Print followed by os.Exit(1)
func (l *Logger) Fatal(a ...interface{}) {
	l.log.Output(2, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf is a Printf followed by os.Exit(1)
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.log.Output(2, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalln is a Println followed by os.Exit(1)
func (l *Logger) Fatalln(a ...interface{}) {
	l.log.Output(2, fmt.Sprintln(a...))
	os.Exit(1)
}

// Panic is a Print followed by panic
func (l *Logger) Panic(a ...interface{}) {
	l.log.Output(2, fmt.Sprint(a...))
	panic(fmt.Sprint(a...))
}

// Panicf is a Printf followed by panic
func (l *Logger) Panicf(format string, a ...interface{}) {
	l.log.Output(2, fmt.Sprintf(format, a...))
	panic(fmt.Sprintf(format, a...))
}

// Panicln is a Println followed by panic
func (l *Logger) Panicln(a ...interface{}) {
	l.log.Output(2, fmt.Sprintln(a...))
	panic(fmt.Sprintln(a...))
}

// Debug is a Print which prints to the logger only when debugging is enabled
func (l *Logger) Debug(a ...interface{}) {
	if l.debug != nil {
		l.debug.Output(2, fmt.Sprint(a...))
	}
}

// Debugf is a Printf which prints to the logger only when debugging is enabled
func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.debug != nil {
		l.debug.Output(2, fmt.Sprintf(format, a...))
	}
}

// Debugln is a Println which prints to the logger only when debugging is enabled
func (l *Logger) Debugln(a ...interface{}) {
	if l.debug != nil {
		l.debug.Output(2, fmt.Sprintln(a...))
	}
}

// DebugDepth is similar to Debug but allows control over the scope of tracing.
// E.g. :
//  10 apiError("error message")
//  11
//  12 func apiError(s string){
//  13 	logger.DebugDepth(2, "API Error:", s)
//  14 }
// Will print: 2016/07/19 17:34:10 [debug] main.go:10: API Error: error message
func (l *Logger) DebugDepth(calldepth int, a ...interface{}) {
	if l.debug != nil {
		l.debug.Output(calldepth+1, fmt.Sprint(a...))
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) error {
	if l.debug != nil {
		l.debug.Output(2, fmt.Sprintf(format, a...))
	}
	return fmt.Errorf(format, a...)
}
