package main

import (
	"context"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"order created",
		slog.Int("id", 23123),
		slog.Group("order",
			slog.Time("order_date", time.Now()),
			slog.String("shipping_address", "123 Main St, City, Country"),
			slog.String("payment_method", "Credit Card"),
			slog.Float64("total_amount", 49.99),
		),
	)
}
