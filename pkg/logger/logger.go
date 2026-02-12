package logger

import (
	"log/slog"
	"os"
	"ps-go-fiber/blog/config"
)

func NewLogger(config *config.LogConfig) *slog.Logger {
    opts := &slog.HandlerOptions{
        Level: slog.Level(config.Level), // <-- level comes from config here
    }

    var handler slog.Handler
    switch config.Format {
    case "json":
        handler = slog.NewJSONHandler(os.Stdout, opts)
    default:
        handler = slog.NewTextHandler(os.Stdout, opts)
    }

    return slog.New(handler)
}
