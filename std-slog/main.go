package main

import (
	"golang.org/x/exp/slog"
)

func main() {
	slog.Debug("Debug message example")
	slog.Info("Info message example")
	slog.Warn("Warning message example")
	slog.Error("Error message example")
}
