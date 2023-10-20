package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/tests"
	"strings"
	"testing"
)

// TestErrorEncoder uses update program version endpoint to test error encoder.
// Is supposed that all endpoints use error encoder middleware by default.
// And if error encoder works for one endpoint it means that it works for all endpoints.
func TestErrorEncoder(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)
	versionID := version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")

	// Test operation
	req, err := http.NewRequest(
		"PUT",
		test.Addr+"/versions/"+versionID.String(),
		strings.NewReader(`{"name":"new-name"}`),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.TODO())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	// Test assertions
	require.Equal(t,
		`{"error":"version not found","code":"NOT_FOUND"}`+"\n",
		string(data),
	)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
