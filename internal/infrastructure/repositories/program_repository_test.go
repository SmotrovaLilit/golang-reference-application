package repositories_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"reference-application/internal/domain/program"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

var testProgramID = program.MustNewID("6f995ea2-3144-4499-b69b-09bd8635404f")

func TestProgramRepository_Save(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	programRepository := repositories.NewProgramRepository(test.DB)

	test.ExpectBegin()
	test.ExpectExec(`UPDATE "programs"`).
		WithArgs("ANDROID", testProgramID.String()).
		WillReturnResult(sqlmock.NewResult(0, 0))
	test.ExpectCommit()

	test.ExpectBegin()
	test.ExpectExec(`INSERT INTO "programs"`).
		WithArgs(testProgramID.String(), "ANDROID").
		WillReturnResult(sqlmock.NewResult(0, 1))
	test.ExpectCommit()

	_program := program.NewProgram(
		testProgramID,
		program.AndroidPlatformCode,
	)
	programRepository.Save(context.Background(), _program)

	require.NoError(t, test.ExpectationsWereMet())
}

func TestProgramRepository_Save_Error(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	programRepository := repositories.NewProgramRepository(test.DB)

	test.ExpectBegin()
	test.ExpectExec(`UPDATE "programs"`).
		WithArgs("ANDROID", testProgramID.String()).
		WillReturnError(errors.New("error"))
	test.ExpectRollback()

	_program := program.NewProgram(
		testProgramID,
		program.AndroidPlatformCode,
	)

	require.Panics(t, func() {
		programRepository.Save(context.Background(), _program)
	})

	require.NoError(t, test.ExpectationsWereMet())
}

func TestProgramRepository_FindByID(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	programRepository := repositories.NewProgramRepository(test.DB)

	test.ExpectQuery(`SELECT \* FROM "programs"`).
		WithArgs(testProgramID.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "platform_code"}).
			AddRow(testProgramID.String(), "ANDROID"))

	_program := programRepository.FindByID(context.Background(), testProgramID)
	require.Equal(t, testProgramID, _program.ID())
	require.Equal(t, program.AndroidPlatformCode, _program.PlatformCode())

	require.NoError(t, test.ExpectationsWereMet())
}

func TestProgramRepository_FindByID_Error(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	programRepository := repositories.NewProgramRepository(test.DB)

	test.ExpectQuery(`SELECT \* FROM "programs"`).
		WithArgs(testProgramID.String()).
		WillReturnError(errors.New("error"))

	require.Panics(t, func() {
		programRepository.FindByID(context.Background(), testProgramID)
	})

	require.NoError(t, test.ExpectationsWereMet())
}
