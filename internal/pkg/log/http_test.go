package log

import (
	"bytes"
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
)

func TestWithHTTPRequestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestMethod, "GET")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestURI, "/test")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestPath, "/path")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestProto, "HTTP/1.1")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestHost, "localhost:8080")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestRemoteAddr, "1")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestXForwardedFor, "2")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestXForwardedProto, "3")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestAuthorization, "4")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestReferer, "5")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestUserAgent, "6")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestXRequestID, "7")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestAccept, "8")

	var writer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&writer, nil))
	gotLogger := WithHTTPRequestContext(ctx, logger)
	gotLogger.Info("test")

	require.Contains(t, writer.String(), "request.method=GET")
	require.Contains(t, writer.String(), "request.uri=/test")
	require.Contains(t, writer.String(), "request.path=/path")
	require.Contains(t, writer.String(), "request.proto=HTTP/1.1")
	require.Contains(t, writer.String(), "request.host=localhost:8080")
	require.Contains(t, writer.String(), "request.remoteAddr=1")
	require.Contains(t, writer.String(), "request.XForwardedFor=2")
	require.Contains(t, writer.String(), "request.XForwardedProto=3")
	require.Contains(t, writer.String(), "request.authorization=4")
	require.Contains(t, writer.String(), "request.referer=5")
	require.Contains(t, writer.String(), "request.userAgent=6")
	require.Contains(t, writer.String(), "request.XRequestID=7")
	require.Contains(t, writer.String(), "request.accept=8")
}
