package main

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("file_path", "/path/to/file.txt"),
		slog.Int("status", 200),
		slog.String(
			"shipping_address",
			"123 Main St, City, Country",
		),
	)
}
