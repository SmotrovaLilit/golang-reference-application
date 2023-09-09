package repositories_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/optional"
	"reference-application/internal/pkg/tests"
	"testing"
)

var testVersionID = version.MustNewID("1f111ea1-3144-4499-b69b-01bd1111111f")
var testVersionName = version.MustNewName("smart-calculator")

func TestVersionRepository_Save(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectBegin()
	test.ExpectExec(`UPDATE "versions"`).
		WithArgs(
			testVersionName.String(),
			testProgramID.String(),
			version.DraftStatus,
			"", // Description
			"", // Number
			testVersionID.String(),
		).
		WillReturnResult(sqlmock.NewResult(0, 0))
	test.ExpectCommit()

	test.ExpectBegin()
	test.ExpectExec(`INSERT INTO "versions"`).
		WithArgs(
			testVersionID.String(),
			testVersionName.String(),
			testProgramID.String(),
			version.DraftStatus,
			"", // Description
			"", // Number
		).
		WillReturnResult(sqlmock.NewResult(0, 1))
	test.ExpectCommit()

	_version := version.NewVersion(
		testVersionID,
		testVersionName,
		testProgramID,
	)
	versionRepository.Save(context.Background(), _version)
	require.NoError(t, test.ExpectationsWereMet())
}

func TestVersionRepository_SaveWithAllFields(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectBegin()
	test.ExpectExec(`UPDATE "versions"`).
		WithArgs(
			testVersionName.String(),
			testProgramID.String(),
			version.DraftStatus,
			"so long description",
			"1.0.0",
			testVersionID.String(),
		).
		WillReturnResult(sqlmock.NewResult(0, 0))
	test.ExpectCommit()

	test.ExpectBegin()
	test.ExpectExec(`INSERT INTO "versions"`).
		WithArgs(
			testVersionID.String(),
			testVersionName.String(),
			testProgramID.String(),
			version.DraftStatus,
			"so long description",
			"1.0.0",
		).
		WillReturnResult(sqlmock.NewResult(0, 1))
	test.ExpectCommit()

	_version := version.NewExistingVersion(
		testVersionID,
		testVersionName,
		testProgramID,
		version.DraftStatus,
		optional.Of[version.Description](version.MustNewDescription("so long description")),
		optional.Of[version.Number](version.MustNewNumber("1.0.0")),
	)
	versionRepository.Save(context.Background(), _version)
	require.NoError(t, test.ExpectationsWereMet())
}

func TestVersionRepository_Save_Error(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectBegin()
	test.ExpectExec(`UPDATE "versions"`).
		WithArgs(testVersionName.String(), testProgramID.String(), version.DraftStatus, "", "", testVersionID.String()).
		WillReturnError(errors.New("error"))
	test.ExpectRollback()

	_version := version.NewVersion(
		testVersionID,
		testVersionName,
		testProgramID,
	)

	require.Panics(t, func() {
		versionRepository.Save(context.Background(), _version)
	})

	require.NoError(t, test.ExpectationsWereMet())
}

func TestVersionRepository_FindByID(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectQuery(`SELECT \* FROM "versions"`).
		WithArgs(testVersionID.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "program_id", "status", "description", "number"}).
			AddRow(testVersionID.String(), testVersionName.String(), testProgramID.String(), "DRAFT", "", ""))

	_version := versionRepository.FindByID(context.Background(), testVersionID)
	require.Equal(t, testVersionID, _version.ID())
	require.Equal(t, testVersionName, _version.Name())
	require.Equal(t, testProgramID, _version.ProgramID())
	require.Equal(t, version.DraftStatus, _version.Status())
	require.True(t, _version.Description().IsEmpty())
	require.True(t, _version.Number().IsEmpty())

	require.NoError(t, test.ExpectationsWereMet())
}

func TestVersionRepository_FindByIDWithAllFields(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectQuery(`SELECT \* FROM "versions"`).
		WithArgs(testVersionID.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "program_id", "status", "description", "number"}).
			AddRow(
				testVersionID.String(),
				testVersionName.String(),
				testProgramID.String(),
				"DRAFT",
				"so long description",
				"1.0.0",
			))

	_version := versionRepository.FindByID(context.Background(), testVersionID)
	require.Equal(t, testVersionID, _version.ID())
	require.Equal(t, testVersionName, _version.Name())
	require.Equal(t, testProgramID, _version.ProgramID())
	require.Equal(t, version.DraftStatus, _version.Status())
	require.True(t, _version.Description().IsPresent())
	require.Equal(t, "so long description", _version.Description().Value().String())
	require.True(t, _version.Number().IsPresent())
	require.Equal(t, "1.0.0", _version.Number().Value().String())

	require.NoError(t, test.ExpectationsWereMet())
}

func TestVersionRepository_FindByID_Error(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectQuery(`SELECT \* FROM "versions"`).
		WithArgs(testVersionID.String()).
		WillReturnError(errors.New("error"))

	require.Panics(t, func() {
		versionRepository.FindByID(context.Background(), testVersionID)
	})

	require.NoError(t, test.ExpectationsWereMet())
}

func TestVersionRepository_FindByID_NotFound(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	versionRepository := repositories.NewVersionRepository(test.DB)

	test.ExpectQuery(`SELECT \* FROM "versions"`).
		WithArgs(testVersionID.String()).
		WillReturnError(gorm.ErrRecordNotFound)

	v := versionRepository.FindByID(context.Background(), testVersionID)
	require.Nil(t, v)

	require.NoError(t, test.ExpectationsWereMet())
}
