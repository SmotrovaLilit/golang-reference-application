package repositories_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	interfaces "reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestUnitOfWork_Do(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)

	testProgramID1 := program.MustNewID("6f995ea2-3144-4499-b69b-09bd8635404f")
	testProgramID2 := program.MustNewID("01be4f16-11dd-4105-b6e9-3ab9bef7ead3")
	unitOfWork := repositories.NewUnitOfWork(test.DB)

	test.ExpectBegin()
	test.ExpectExec(`UPDATE "programs"`).
		WithArgs("ANDROID", testProgramID1.String()).
		WillReturnResult(sqlmock.NewResult(0, 0))
	test.ExpectExec(`INSERT INTO "programs"`).
		WithArgs(testProgramID1.String(), "ANDROID").
		WillReturnResult(sqlmock.NewResult(0, 1))
	test.ExpectExec(`UPDATE "programs"`).
		WithArgs("ANDROID", testProgramID2.String()).
		WillReturnResult(sqlmock.NewResult(0, 0))
	test.ExpectExec(`INSERT INTO "programs"`).
		WithArgs(testProgramID2.String(), "ANDROID").
		WillReturnResult(sqlmock.NewResult(0, 1))
	test.ExpectCommit()

	unitOfWork.Do(context.TODO(), func(store interfaces.UnitOfWorkStore) {
		store.ProgramRepository().Save(context.Background(), program.NewProgram(
			testProgramID1,
			program.AndroidPlatformCode,
		))
		store.ProgramRepository().Save(context.Background(), program.NewProgram(
			testProgramID2,
			program.AndroidPlatformCode,
		))
	})

	require.NoError(t, test.ExpectationsWereMet())
}

func TestUnitOfWork_Do_Panic(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)

	unitOfWork := repositories.NewUnitOfWork(test.DB)

	test.ExpectBegin()
	test.ExpectRollback()

	require.Panics(t, func() {
		unitOfWork.Do(context.TODO(), func(store interfaces.UnitOfWorkStore) {
			panic(errors.New("some error"))
		})
	})

	require.NoError(t, test.ExpectationsWereMet())
}
