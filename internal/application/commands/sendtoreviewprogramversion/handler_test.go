package sendtoreviewprogramversion_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	"reference-application/internal/application/commands/sendtoreviewprogramversion"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := sendtoreviewprogramversion.Handler{
		Repository: versionRepository,
		Logger:     slog.Default(),
	}
	endpoint := sendtoreviewprogramversion.NewEndpoint(handler)

	// Prepare test data
	existingVersion := dbTest.PrepareDraftVersionReadyToReview(t)

	// Tested operation
	cmd := sendtoreviewprogramversion.NewCommand(
		existingVersion.ID(),
	)
	_, err := endpoint(context.TODO(), cmd)

	// Test assertions
	require.NoError(t, err)
	_version := versionRepository.FindByID(context.Background(), existingVersion.ID())
	require.NotNil(t, _version)
	require.True(t, _version.Status().IsOnReview())
}

func TestHandler_HandleVersionNotFound(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := sendtoreviewprogramversion.Handler{
		Repository: versionRepository,
		Logger:     slog.Default(),
	}
	endpoint := sendtoreviewprogramversion.NewEndpoint(handler)

	// Prepare test data
	_version, _ := tests.NewPreparedToReviewVersion()
	versionID := _version.ID()

	// Tested operation
	cmd := sendtoreviewprogramversion.NewCommand(
		versionID,
	)
	_, err := endpoint(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, sharederrors.ErrVersionNotFound)
}

func TestHandler_HandleErrorFromDomain(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := sendtoreviewprogramversion.Handler{
		Repository: versionRepository,
		Logger:     slog.Default(),
	}
	endpoint := sendtoreviewprogramversion.NewEndpoint(handler)

	// Prepare test data
	existingVersion := dbTest.PrepareVersionOnReview(t)

	// Tested operation
	cmd := sendtoreviewprogramversion.NewCommand(
		existingVersion.ID(),
	)
	_, err := endpoint(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, version.ErrInvalidStatusToSendToReview)
}

func TestHandler_Endpoint(t *testing.T) {
	t.Run("endpoint should return resource name and action", func(t *testing.T) {
		endpoint := sendtoreviewprogramversion.NewEndpoint(sendtoreviewprogramversion.Handler{})
		require.Equal(t, "programVersion", endpoint.ResourceName())
		require.Equal(t, "sendToReview", endpoint.ResourceAction())
	})
}
