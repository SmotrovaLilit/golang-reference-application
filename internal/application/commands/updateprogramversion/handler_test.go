package updateprogramversion_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/optional"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := updateprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	versionNewName := version.MustNewName("new-name")
	newDescription := optional.Of[version.Description](version.MustNewDescription("new-description"))
	newNumber := optional.Of[version.Number](version.MustNewNumber("1.0.1"))
	existingVersion := dbTest.PrepareDraftVersion(t)

	// Tested operation
	cmd := updateprogramversion.NewCommand(
		existingVersion.ID(),
		versionNewName,
		newDescription,
		newNumber,
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.NoError(t, err)
	_version := versionRepository.FindByID(context.Background(), existingVersion.ID())
	require.NotNil(t, _version)
	require.Equal(t, versionNewName, _version.Name())
	require.Equal(t, versionNewName, _version.Name())
	require.Equal(t, newDescription, _version.Description())
	require.Equal(t, newNumber, _version.Number())
}

func TestHandler_HandleVersionNotFound(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := updateprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	_version, _ := tests.NewDraftVersion()
	versionID := _version.ID()
	versionNewName := version.MustNewName("new-name")

	// Tested operation
	cmd := updateprogramversion.NewCommand(
		versionID,
		versionNewName,
		optional.Empty[version.Description](),
		optional.Empty[version.Number](),
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, sharederrors.ErrVersionNotFound)
}

func TestHandler_HandleErrorFromDomainUpdateVersion(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := updateprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	existingVersion := dbTest.PrepareVersionOnReview(t)
	versionNewName := version.MustNewName("new-name")
	versionOldName := existingVersion.Name()

	// Tested operation
	cmd := updateprogramversion.NewCommand(
		existingVersion.ID(),
		versionNewName,
		optional.Empty[version.Description](),
		optional.Empty[version.Number](),
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, version.ErrUpdateVersionStatus)
	_version := versionRepository.FindByID(context.Background(), existingVersion.ID())
	require.NotNil(t, _version)
	require.Equal(t, versionOldName, _version.Name())
}
