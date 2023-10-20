package log

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"
)

// WithContext adds context to logger and returns new Logger.
func WithContext(ctx context.Context, logger *slog.Logger) *slog.Logger {
	logger = WithHTTPRequestContext(ctx, logger)
	logger = WithResourceContext(ctx, logger)
	return logger
}

// WithApplicationInfo adds application info to logger and returns new Logger.
func WithApplicationInfo(logger *slog.Logger, applicationName string) *slog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()
	return logger.With(
		slog.Group("application",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
			slog.String("version", buildInfo.Main.Version),
			slog.String("name", applicationName),
		),
	)
}
