package main

// START OMIT
import (
	"log/slog" // HL12
)

func main() {
	slog.Debug("Debug message example")
	slog.Info("Info message example")
	slog.Warn("Warning message example")
	slog.Error("Error message example")
}

// END OMIT
