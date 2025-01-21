package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func StartLogger() {
	once.Do(func() {
		opts := &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}

		handler := slog.NewJSONHandler(os.Stderr, opts)
		logger = slog.New(handler)
		slog.SetDefault(logger)
	})
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	logger.Error(msg, args...)
}
