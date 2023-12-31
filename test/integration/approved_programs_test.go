package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"os"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/tests"
	"testing"
)

func TestApprovedProgramsHandler(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)
	prepareData(t, test)

	// Test operation
	req, err := http.NewRequest(
		"GET",
		test.Addr+"/store/programs",
		nil,
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.TODO())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})

	// Test assertions
	gotBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, getExpectedBody(t), string(gotBody))
}

func getExpectedBody(t *testing.T) string {
	expectedBody, err := os.ReadFile("testdata/approved_programs_response.json")
	require.NoError(t, err)
	return string(expectedBody)
}

func prepareData(t *testing.T, test tests.IntegrationTest) {
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
	test.SavePrograms(t, []program.Program{
		approvedProgram1,
		approvedProgram2,
		programWithoutVersions,
		programWithoutApprovedVersions,
	})
	test.SaveVersions(t, []version.Version{
		draftVersion1,
		draftVersion2,
		approvedVersion1,
		approvedVersion2,
		draftVersion3,
	})
}
