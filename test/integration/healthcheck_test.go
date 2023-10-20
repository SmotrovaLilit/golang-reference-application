package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/pkg/tests"

	"net/http"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)

	request, err := http.NewRequest("GET", test.Addr+"/health", nil)
	request = request.WithContext(context.TODO())
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHealthCheckFailed(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)

	request, err := http.NewRequest("GET", test.Addr+"/health", nil)
	request = request.WithContext(context.TODO())
	require.NoError(t, err)
	test.TestWithDatabase.TerminateDatabase(t)
	time.Sleep(2 * time.Second)
	resp, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})
	require.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
}
