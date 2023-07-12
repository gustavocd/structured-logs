package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, &opts)

	logger := slog.New(handler)

	logger.Debug("Debug message example")
	logger.Info("Info message example")
	logger.Warn("Warning message example")
	logger.Error("Error message example")
}
