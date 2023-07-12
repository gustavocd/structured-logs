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
		"file_path", // this key has no value
	)
}
