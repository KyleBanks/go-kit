package log

import "testing"

// NOTE: Currently we just call the methods and ensure nothing blows up.
// Don't have a way to actually test that it's logging anything.

func TestLogger_Print(t *testing.T) {
	Logger.Print("a", "b", "c")
	Logger.Print()
}

func TestInfo(t *testing.T) {
	Info("A", "B", "C")
	Info()
}

func TestError(t *testing.T) {
	Error("a", "b", "c")
	Error()
}

func TestPrintStack(t *testing.T) {
	PrintStack()
}
