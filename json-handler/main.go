package main

import (
	"os"

	"log/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Debug("Debug message example")
	logger.Info("Info message example")
	logger.Warn("Warning message example")
	logger.Error("Error message example")
}
