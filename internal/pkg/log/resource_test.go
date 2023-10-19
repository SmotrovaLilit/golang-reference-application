package log

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"

	"reference-application/internal/pkg/resource"
)

func TestWithResourceContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, resource.ContextKeyResourceID, "1")
	ctx = context.WithValue(ctx, resource.ContextKeyResourceType, "2")
	ctx = context.WithValue(ctx, resource.ContextKeyResourceAction, "3")

	var writer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&writer, nil))
	gotLogger := WithResourceContext(ctx, logger)
	gotLogger.Info("test")

	require.Contains(t, writer.String(), "resource.id=1")
	require.Contains(t, writer.String(), "resource.type=2")
	require.Contains(t, writer.String(), "resource.action=3")
}
