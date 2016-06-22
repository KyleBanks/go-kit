// Simple logging service to print to stdout/stderr
package log

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

type logger struct{}

var Logger logger

// Println is a wrapper for the Info log method.
//
// This is used for passing the Logger around as a LogWriter interface.
func (l logger) Print(a ...interface{}) {
	Info(a)
}

// Info outputs to stdout.
func Info(a ...interface{}) {
	fmt.Println(time.Now(), a)
}

// Infof outputs a formatted string to stdout.
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Error outputs to stderr.
func Error(a ...interface{}) {
	fmt.Fprintln(os.Stderr, time.Now(), a)
}

// Errorf outputs a formatted error to stderr.
func Errorf(format string, a ...interface{}) {
	Error(fmt.Sprintf(format, a...))
}

// PrintStack outputs the current go routine's stack trace.
func PrintStack() {
	debug.PrintStack()
}
