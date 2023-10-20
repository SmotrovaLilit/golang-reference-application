package integration

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"reference-application/internal/pkg/tests"
	"testing"
	"time"
)

// TestPanicRecovery covers only cases when panic is happened in application layer.
// Test for approve program version is used.
// Is supposed that all endpoints use panic recovery middleware by default.
// And if panic recovery works for one endpoint it means that it works for all endpoints.
func TestPanicRecovery(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)

	// Test operation
	req, err := http.NewRequest(
		"PUT",
		test.Addr+fmt.Sprintf("/versions/%s/approve", "ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
		nil,
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.TODO())

	test.TerminateDatabase(t)
	time.Sleep(2 * time.Second)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, `{"error":"Internal Server Error","code":"INTERNAL_SERVER_ERROR"}`+"\n", string(data))
}
