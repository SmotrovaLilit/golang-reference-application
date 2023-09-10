package integration

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestSendToReviewProgramVersionHandler(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)
	existingVersion := test.PrepareDraftVersionReadyToReview(t)

	// Test operation
	req, err := http.NewRequest(
		"PUT",
		test.Addr+fmt.Sprintf("/versions/%s/sendToReview", existingVersion.ID().String()),
		nil,
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.TODO())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})

	// Test assertions
	versionRepository := repositories.NewVersionRepository(test.DB)
	_version := versionRepository.FindByID(context.TODO(), existingVersion.ID())
	require.NotNil(t, _version)
	require.True(t, _version.Status().IsOnReview())
}
