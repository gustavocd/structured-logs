package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"file_path", "/path/to/file.txt",
		"status", 200,
	)
}
