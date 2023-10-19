package log

import (
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"log/slog"
)

// WithHTTPRequestContext returns a new logger with the HTTP request context.
func WithHTTPRequestContext(ctx context.Context, logger *slog.Logger) *slog.Logger {
	var attrs []any
	for name, key := range map[string]interface{}{
		"method":          kithttp.ContextKeyRequestMethod,
		"uri":             kithttp.ContextKeyRequestURI,
		"path":            kithttp.ContextKeyRequestPath,
		"proto":           kithttp.ContextKeyRequestProto,
		"host":            kithttp.ContextKeyRequestHost,
		"remoteAddr":      kithttp.ContextKeyRequestRemoteAddr,
		"XForwardedFor":   kithttp.ContextKeyRequestXForwardedFor,
		"XForwardedProto": kithttp.ContextKeyRequestXForwardedProto,
		"authorization":   kithttp.ContextKeyRequestAuthorization,
		"referer":         kithttp.ContextKeyRequestReferer,
		"userAgent":       kithttp.ContextKeyRequestUserAgent,
		"XRequestID":      kithttp.ContextKeyRequestXRequestID,
		"accept":          kithttp.ContextKeyRequestAccept,
	} {
		if value, ok := ctx.Value(key).(string); ok {
			attrs = append(attrs, slog.String(name, value))
		}
	}
	return logger.With(slog.Group("request", attrs...))
}
