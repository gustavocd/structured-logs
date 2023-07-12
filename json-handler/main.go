package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Debug("Debug message example")
	logger.Info("Info message example")
	logger.Warn("Warning message example")
	logger.Error("Error message example")
}
