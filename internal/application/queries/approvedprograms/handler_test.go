package approvedprograms_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/queries/approvedprograms"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/readmodels"
	"reference-application/internal/pkg/pager"
	"reference-application/internal/pkg/tests"
	"testing"
)

//nolint:funlen
func TestHandler_Handle(t *testing.T) {
	dbTest := tests.PrepareTestWithDatabase(t)
	readModel := readmodels.NewApprovedProgramsReadModel(dbTest.DB)
	handler := approvedprograms.Handler{
		ReadModel: readModel,
	}
	endpoint := approvedprograms.NewEndpoint(handler)

	// Prepare data.
	approvedProgram1 := tests.NewProgram(
		program.MustNewID("39bc0bbb-a657-4208-b672-d79e7977328f"),
	)
	draftVersion1 := tests.NewDraftVersionInProgram(
		approvedProgram1,
		version.MustNewID("18a0a8ce-4eb1-4588-9ba2-eb111262f236"),
	)
	approvedVersion1 := tests.NewApprovedVersionInProgram(
		approvedProgram1,
		version.MustNewID("45d14653-5228-4a0e-ada5-aa723a85cc35"),
	)

	approvedProgram2 := tests.NewProgram(
		program.MustNewID("7b83b402-7681-486e-aaa8-816babebd81c"),
	)
	draftVersion2 := tests.NewDraftVersionInProgram(
		approvedProgram2,
		version.MustNewID("44b02ed9-fbc5-44b7-b9e8-05ab02c6291e"),
	)
	approvedVersion2 := tests.NewApprovedVersionInProgram(
		approvedProgram2,
		version.MustNewID("d7fcaf7d-16d6-4219-829f-03a91f32a202"),
	)

	programWithoutVersions := tests.NewProgram(
		program.MustNewID("00f1c380-c88f-4664-a07b-94a82b5205d4"),
	)
	programWithoutApprovedVersions := tests.NewProgram(
		program.MustNewID("cb0c011a-5b60-4b0d-8234-34d19f80a21e"),
	)
	draftVersion3 := tests.NewOnReviewVersionInProgram(
		programWithoutApprovedVersions,
		version.MustNewID("1a28ff2d-172e-4e66-9eae-dc4a746acdf7"),
	)

	dbTest.SavePrograms(t, []program.Program{
		approvedProgram1,
		approvedProgram2,
		programWithoutVersions,
		programWithoutApprovedVersions,
	})
	dbTest.SaveVersions(t, []version.Version{
		draftVersion1,
		draftVersion2,
		approvedVersion1,
		approvedVersion2,
		draftVersion3,
	})

	// Tested operation
	query := approvedprograms.NewQuery(pager.Default)
	resp, err := endpoint(context.TODO(), query)
	require.NoError(t, err)
	result := resp.(approvedprograms.Result)

	// Test assertions
	require.Len(t, result, 2)
	require.Equal(t, approvedProgram1.ID().String(), result[0].ID)
	require.Equal(t, approvedProgram1.PlatformCode().String(), result[0].PlatformCode)
	require.Equal(t, approvedVersion1.ID().String(), result[0].Version.ID)
	require.Equal(t, approvedVersion1.Name().String(), result[0].Version.Name)
	require.Equal(t, approvedVersion1.Number().Value().String(), result[0].Version.Number)
	require.Equal(t, approvedVersion1.Description().Value().String(), result[0].Version.Description)
	require.Equal(t, approvedVersion1.Status().String(), result[0].Version.Status)

	require.Equal(t, approvedProgram2.ID().String(), result[1].ID)
	require.Equal(t, approvedProgram2.PlatformCode().String(), result[1].PlatformCode)
	require.Equal(t, approvedVersion2.ID().String(), result[1].Version.ID)
	require.Equal(t, approvedVersion2.Name().String(), result[1].Version.Name)
	require.Equal(t, approvedVersion2.Number().Value().String(), result[1].Version.Number)
	require.Equal(t, approvedVersion2.Description().Value().String(), result[1].Version.Description)
	require.Equal(t, approvedVersion2.Status().String(), result[1].Version.Status)
}
