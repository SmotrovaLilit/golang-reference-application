package updateprogramversion_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	programRepository := repositories.NewProgramRepository(dbTest.DB)
	handler := updateprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	versionID := version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	versionNewName := version.MustNewName("new-name")
	programID := program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")
	programRepository.Save(context.TODO(), program.NewProgram(
		programID,
		program.AndroidPlatformCode,
	))
	versionRepository.Save(context.TODO(), version.NewVersion(
		versionID,
		version.MustNewName("smart-calculator"),
		programID,
	))

	// Tested operation
	cmd := updateprogramversion.NewCommand(
		versionID,
		versionNewName,
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.NoError(t, err)
	_version := versionRepository.FindByID(context.Background(), versionID)
	require.NotNil(t, _version)
	require.Equal(t, versionNewName, _version.Name())
}

func TestHandler_HandleVersionNotFound(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	handler := updateprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	versionID := version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	versionNewName := version.MustNewName("new-name")

	// Tested operation
	cmd := updateprogramversion.NewCommand(
		versionID,
		versionNewName,
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, updateprogramversion.ErrVersionNotFound)
}

func TestHandler_HandleErrorFromDomainUpdateVersion(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)
	programRepository := repositories.NewProgramRepository(dbTest.DB)
	handler := updateprogramversion.Handler{
		Repository: versionRepository,
	}

	// Prepare test data
	versionID := version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	versionNewName := version.MustNewName("new-name")
	versionOldName := version.MustNewName("smart-calculator")
	programID := program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")
	programRepository.Save(context.TODO(), program.NewProgram(
		programID,
		program.AndroidPlatformCode,
	))
	versionRepository.Save(context.TODO(), version.NewExistingVersion(
		versionID,
		versionOldName,
		programID,
		version.OnReviewStatus,
	))

	// Tested operation
	cmd := updateprogramversion.NewCommand(
		versionID,
		versionNewName,
	)
	err := handler.Handle(context.TODO(), cmd)

	// Test assertions
	require.ErrorIs(t, err, version.ErrUpdateStatus)
	_version := versionRepository.FindByID(context.Background(), versionID)
	require.NotNil(t, _version)
	require.Equal(t, versionOldName, _version.Name())
}
