package createprogram_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/domain/program"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	programRepository := repositories.NewProgramRepository(dbTest.DB)
	handler := createprogram.Handler{
		Repository: programRepository,
	}

	cmd := createprogram.NewCommand(
		program.MustNewID("35A530CF-91F3-49DC-BB6D-AC423563541C"),
		program.AndroidPlatformCode,
	)
	handler.Handle(context.Background(), cmd)

	_program := programRepository.FindByID(context.Background(), cmd.ID)
	require.NotNil(t, _program)
	require.Equal(t, cmd.ID, _program.ID())
	require.Equal(t, cmd.PlatformCode, _program.PlatformCode())
}
