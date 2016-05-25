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

// Error outputs to stderr.
func Error(a ...interface{}) {
	fmt.Fprintln(os.Stderr, time.Now(), a)
}

// PrintStack outputs the current go routine's stack trace.
func PrintStack() {
	debug.PrintStack()
}
