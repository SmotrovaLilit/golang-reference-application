package declineprogramversion_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	"reference-application/internal/application/commands/declineprogramversion"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := declineprogramversion.Handler{
		Repository: versionRepository,
		Logger:     slog.Default(),
	}
	endpoint := declineprogramversion.NewEndpoint(handler)

	// Prepare test data
	existingVersion := dbTest.PrepareVersionOnReview(t)

	// Tested operation
	cmd := declineprogramversion.NewCommand(
		existingVersion.ID(),
	)
	_, err := endpoint(context.TODO(), cmd)

	// Test assertions
	require.NoError(t, err)
	_version := versionRepository.FindByID(context.Background(), existingVersion.ID())
	require.NotNil(t, _version)
	require.True(t, _version.Status().IsDeclined())
}

func TestHandler_HandleVersionNotFound(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := declineprogramversion.Handler{
		Repository: versionRepository,
		Logger:     slog.Default(),
	}
	endpoint := declineprogramversion.NewEndpoint(handler)

	// Prepare test data
	_version, _ := tests.NewOnReviewVersion()
	versionID := _version.ID()

	// Tested operation
	cmd := declineprogramversion.NewCommand(
		versionID,
	)
	_, err := endpoint(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, sharederrors.ErrVersionNotFound)
}

func TestHandler_HandleErrorFromDomain(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := declineprogramversion.Handler{
		Repository: versionRepository,
		Logger:     slog.Default(),
	}
	endpoint := declineprogramversion.NewEndpoint(handler)

	// Prepare test data
	existingVersion := dbTest.PrepareDraftVersion(t)

	// Tested operation
	cmd := declineprogramversion.NewCommand(
		existingVersion.ID(),
	)
	_, err := endpoint(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, version.ErrInvalidStatusToDecline)
}

func TestHandler_Endpoint(t *testing.T) {
	t.Run("endpoint should return resource name and action", func(t *testing.T) {
		endpoint := declineprogramversion.NewEndpoint(declineprogramversion.Handler{})
		require.Equal(t, "programVersion", endpoint.ResourceName())
		require.Equal(t, "decline", endpoint.ResourceAction())
	})
}
