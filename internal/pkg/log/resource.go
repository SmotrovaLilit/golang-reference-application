package log

import (
	"context"
	"log/slog"
	"reference-application/internal/pkg/resource"
)

// WithResourceContext adds resource context to logger.
func WithResourceContext(ctx context.Context, logger *slog.Logger) *slog.Logger {
	var attrs []any
	for name, key := range map[string]interface{}{
		"id":     resource.ContextKeyResourceID,
		"type":   resource.ContextKeyResourceType,
		"action": resource.ContextKeyResourceAction,
	} {
		if value, ok := ctx.Value(key).(string); ok {
			attrs = append(attrs, slog.String(name, value))
		}
	}
	return logger.With(slog.Group("resource", attrs...))
}
