package approveprogramversion_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/approveprogramversion"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := approveprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	existingVersion := dbTest.PrepareVersionOnReview(t)

	// Tested operation
	cmd := approveprogramversion.NewCommand(
		existingVersion.ID(),
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.NoError(t, err)
	_version := versionRepository.FindByID(context.Background(), existingVersion.ID())
	require.NotNil(t, _version)
	require.True(t, _version.Status().IsApproved())
}

func TestHandler_HandleVersionNotFound(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := approveprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	_version, _ := tests.NewOnReviewVersion()
	versionID := _version.ID()

	// Tested operation
	cmd := approveprogramversion.NewCommand(
		versionID,
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, sharederrors.ErrVersionNotFound)
}

func TestHandler_HandleErrorFromDomain(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := approveprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	existingVersion := dbTest.PrepareDraftVersion(t)

	// Tested operation
	cmd := approveprogramversion.NewCommand(
		existingVersion.ID(),
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, version.ErrInvalidStatusToApprove)
}
