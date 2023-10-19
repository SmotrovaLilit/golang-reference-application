package log

import (
	"bytes"
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/require"
	"log/slog"
	"reference-application/internal/pkg/resource"
	"testing"
)

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, resource.ContextKeyResourceID, "1")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestMethod, "GET")

	var writer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&writer, nil))
	gotLogger := WithContext(ctx, logger)
	gotLogger.Info("test")

	require.Contains(t, writer.String(), "resource.id=1")
	require.Contains(t, writer.String(), "request.method=GET")
}

func TestWithApplicationInfo(t *testing.T) {
	writer := bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(&writer, nil))
	gotLogger := WithApplicationInfo(logger, "applicationName")
	gotLogger.Info("test")
	require.Contains(t, writer.String(), "application.name=applicationName")
	require.Contains(t, writer.String(), "application.pid=")
	require.Contains(t, writer.String(), "application.version=")
	require.Contains(t, writer.String(), "application.go_version=go1.")
}
