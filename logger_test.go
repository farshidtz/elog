package elog

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	var debug bool = true
	logger := New("[test] ", &Config{
		TimeFormat:   " ", // remove time for test
		DebugEnabled: &debug,
		DebugPrefix:  "[test-debug] ",
	})
	err := logger.Errorf("error message")
	if err == nil {
		t.Fatal("Error not returned.")
	}
	if err.Error() != "error message" {
		t.Fatalf("Error message is %q instead of %q", err.Error(), "error message")
	}
	// Error with flags
	err = logger.Errorf("error %d", 404)
	if err == nil {
		t.Fatal("Error not returned.")
	}
	if err.Error() != "error 404" {
		t.Fatalf("Error message is %q instead of %q", err.Error(), "error 404")
	}
}

func ExampleErrorf() {
	var debug bool = true
	logger := New("[test] ", &Config{
		TimeFormat:   " ", // remove time for test
		DebugEnabled: &debug,
		DebugPrefix:  "[test-debug] ",
	})
	logger.Errorf("error message")
	// Output: [test-debug] logger_test.go:38: error message
}

func ExamplePrint() {
	logger := New("[test] ", &Config{
		TimeFormat: " ", // remove time for test
	})
	logger.Println("message")
	// Output: [test] message
}

func ExamplePrint_trace() {
	logger := New("[test] ", &Config{
		TimeFormat: " ", // remove time for test
		Trace:      ShortFile,
	})
	logger.Println("message")
	// Output: [test] logger_test.go:55: message
}

func ExampleDebug() {
	var debug bool = true
	logger := New("[test] ", &Config{
		TimeFormat:   " ", // remove time for test
		DebugEnabled: &debug,
		DebugPrefix:  "[test-debug] ",
	})
	logger.Debugln("message")
	// Output: [test-debug] logger_test.go:66: message
}

func ExampleDebug_trace() {
	var debug bool = true
	logger := New("[test] ", &Config{
		TimeFormat:   " ", // remove time for test
		DebugEnabled: &debug,
		DebugPrefix:  "[test-debug] ",
		DebugTrace:   NoTrace,
	})
	logger.Debugln("message")
	// Output: [test-debug] message
}
