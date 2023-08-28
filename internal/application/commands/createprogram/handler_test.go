package createprogram_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
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

	cmd := createprogram.NewCommand(
		program.MustNewID("35A530CF-91F3-49DC-BB6D-AC423563541C"),
		program.AndroidPlatformCode,
		version.MustNewID("11A111CF-91F3-49DC-BB6D-AC4235635411"),
		version.MustNewName("smart-calculator"),
	)
	handler.Handle(context.Background(), cmd)

	_program := programRepository.FindByID(context.Background(), cmd.ID)
	require.NotNil(t, _program)
	require.Equal(t, cmd.ID, _program.ID())
	require.Equal(t, cmd.PlatformCode, _program.PlatformCode())
	_version := versionRepository.FindByID(context.Background(), cmd.VersionID)
	require.NotNil(t, _version)
	require.Equal(t, cmd.VersionID, _version.ID())
	require.Equal(t, cmd.VersionName, _version.Name())
	require.Equal(t, cmd.ID, _version.ProgramID())
}
