package main

import (
	"log"
	"os"

	"log/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Info message example")

	log.Println("Hello from old logger")
}
