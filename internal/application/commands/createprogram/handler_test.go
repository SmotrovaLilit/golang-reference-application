package createprogram_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	programRepository := repositories.NewProgramRepository(dbTest.DB)
	versionRepository := repositories.NewVersionRepository(dbTest.DB)

	handler := createprogram.Handler{
		UnitOfWork: repositories.NewUnitOfWork(dbTest.DB),
	}
	versionInputData, programInputData := tests.NewDraftVersion()
	cmd := createprogram.NewCommand(
		programInputData.ID(),
		programInputData.PlatformCode(),
		versionInputData.ID(),
		versionInputData.Name(),
	)
	handler.Handle(context.Background(), cmd)

	_program := programRepository.FindByID(context.Background(), programInputData.ID())
	require.NotNil(t, _program)
	require.Equal(t, programInputData.ID(), _program.ID())
	require.Equal(t, programInputData.PlatformCode(), _program.PlatformCode())
	_version := versionRepository.FindByID(context.Background(), versionInputData.ID())
	require.NotNil(t, _version)
	require.Equal(t, versionInputData.ID(), _version.ID())
	require.Equal(t, versionInputData.Name(), _version.Name())
	require.Equal(t, programInputData.ID(), _version.ProgramID())
}
