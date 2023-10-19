package approveprogramversion_test

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	"reference-application/internal/application/commands/approveprogramversion"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/resource"
	"reference-application/internal/pkg/tests"
	"testing"
)

// nolint:funlen
func TestHandler_Handle(t *testing.T) {
	t.Run("if input data are valid should approve program version", func(t *testing.T) {
		dbTest := tests.PrepareTestWithDatabase(t)
		versionRepository := repositories.NewVersionRepository(dbTest.DB)
		loggerBuffer := &bytes.Buffer{}
		handler := approveprogramversion.Handler{
			Repository: versionRepository,
			Logger: slog.New(slog.NewTextHandler(loggerBuffer, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
		}
		endpoint := approveprogramversion.NewEndpoint(handler)

		// Prepare test data
		existingVersion := dbTest.PrepareVersionOnReview(t)

		// Tested operation
		cmd := approveprogramversion.NewCommand(
			existingVersion.ID(),
		)
		_, err := endpoint(context.TODO(), cmd)

		// Test assertions
		require.NoError(t, err)
		_version := versionRepository.FindByID(context.Background(), existingVersion.ID())
		require.NotNil(t, _version)
		require.True(t, _version.Status().IsApproved())
		require.Contains(t, loggerBuffer.String(), "program version approved")
	})
	t.Run("if version not found should return error ErrVersionNotFound", func(t *testing.T) {
		dbTest := tests.PrepareTestWithDatabase(t)
		versionRepository := repositories.NewVersionRepository(dbTest.DB)
		loggerBuffer := &bytes.Buffer{}
		handler := approveprogramversion.Handler{
			Repository: versionRepository,
			Logger: slog.New(slog.NewTextHandler(loggerBuffer, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
		}
		endpoint := approveprogramversion.NewEndpoint(handler)

		// Prepare test data
		_version, _ := tests.NewOnReviewVersion()
		versionID := _version.ID()

		// Tested operation
		cmd := approveprogramversion.NewCommand(
			versionID,
		)
		_, err := endpoint(context.TODO(), cmd)

		// Test assertions
		require.ErrorIs(t, err, sharederrors.ErrVersionNotFound)
		require.Contains(t, loggerBuffer.String(),
			"failed to approve program version: program version not found")
	})
	t.Run("if version is not on review should return error ErrVersionIsNotOnReview", func(t *testing.T) {
		dbTest := tests.PrepareTestWithDatabase(t)
		versionRepository := repositories.NewVersionRepository(dbTest.DB)
		loggerBuffer := &bytes.Buffer{}
		handler := approveprogramversion.Handler{
			Repository: versionRepository,
			Logger: slog.New(slog.NewTextHandler(loggerBuffer, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
		}
		endpoint := approveprogramversion.NewEndpoint(handler)

		// Prepare test data
		existingVersion := dbTest.PrepareDraftVersion(t)

		// Tested operation
		cmd := approveprogramversion.NewCommand(
			existingVersion.ID(),
		)
		_, err := endpoint(context.TODO(), cmd)

		// Test assertions
		require.ErrorIs(t, err, version.ErrInvalidStatusToApprove)
		require.Contains(t, loggerBuffer.String(),
			"failed to approve program version: invalid status to approve")
	})
	t.Run("should log resource information", func(t *testing.T) {
		dbTest := tests.PrepareTestWithDatabase(t)
		versionRepository := repositories.NewVersionRepository(dbTest.DB)
		loggerBuffer := &bytes.Buffer{}
		handler := approveprogramversion.Handler{
			Repository: versionRepository,
			Logger: slog.New(slog.NewTextHandler(loggerBuffer, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
		}
		ctx := context.WithValue(context.Background(), resource.ContextKeyResourceAction, "approve")
		err := handler.Handle(ctx, approveprogramversion.NewCommand(version.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")))
		require.Error(t, err)
		require.Contains(t, loggerBuffer.String(), "resource.id=ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")
		require.Contains(t, loggerBuffer.String(), "resource.action=approve")
	})
	t.Run("endpoint should return resource name and action", func(t *testing.T) {
		endpoint := approveprogramversion.NewEndpoint(approveprogramversion.Handler{})
		require.Equal(t, "programVersion", endpoint.ResourceName())
		require.Equal(t, "approve", endpoint.ResourceAction())
	})
}
