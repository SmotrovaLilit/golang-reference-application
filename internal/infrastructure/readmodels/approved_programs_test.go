package readmodels_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"reference-application/internal/infrastructure/readmodels"
	"reference-application/internal/pkg/pager"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestProgramRepository_Query(t *testing.T) {
	test := tests.PrepareTestWithMockedDatabase(t)
	readModel := readmodels.NewApprovedProgramsReadModel(test.DB)

	test.ExpectQuery(`SELECT .* FROM "programs" inner join \(SELECT .* FROM versions WHERE status='APPROVED'\) .* WHERE .* LIMIT 10`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "platform_code", "version_id", "name", "number", "description", "status"}).
			AddRow("6f995ea2-3144-4499-b69b-09bd8635404f", "ANDROID", "11a111cf-91f3-49dc-bb6d-ac4235635411", "name", "1.0.0", "description", "APPROVED"))

	programs := readModel.Query(context.Background(), pager.Default)
	require.Equal(t, "6f995ea2-3144-4499-b69b-09bd8635404f", programs[0].ID)
	require.Equal(t, "ANDROID", programs[0].PlatformCode)
	require.Equal(t, "11a111cf-91f3-49dc-bb6d-ac4235635411", programs[0].Version.ID)
	require.Equal(t, "name", programs[0].Version.Name)
	require.Equal(t, "1.0.0", programs[0].Version.Number)
	require.Equal(t, "description", programs[0].Version.Description)
	require.Equal(t, "APPROVED", programs[0].Version.Status)
	require.NoError(t, test.ExpectationsWereMet())
}
