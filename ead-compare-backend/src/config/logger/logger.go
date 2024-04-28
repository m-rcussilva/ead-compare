package logger

import (
	"log/slog"
	"os"
)

func InitLoger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
