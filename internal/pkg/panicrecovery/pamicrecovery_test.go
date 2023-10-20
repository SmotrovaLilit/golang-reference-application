package panicrecovery

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	xhttp "reference-application/internal/pkg/http"
	"testing"
)

func TestEndpointMiddleware(t *testing.T) {
	endpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		panic("test panic")
		return nil, nil
	}
	loggerBuffer := &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(loggerBuffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	_, err := EndpointMiddleware(endpoint, logger)(context.Background(), nil)
	require.Equal(t, xhttp.ErrInternal, err)
	require.Contains(t, loggerBuffer.String(), "panic recovered: test panic")
	require.Contains(t, loggerBuffer.String(), "level=ERROR")
}
