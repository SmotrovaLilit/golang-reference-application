package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/tests"
	"strings"
	"testing"
)

func TestUpdateProgramVersionHandler(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)
	versionRepository := repositories.NewVersionRepository(test.DB)
	programRepository := repositories.NewProgramRepository(test.DB)
	versionID := version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	newVersionName := version.MustNewName("new-name")
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
	// Test operation
	req, err := http.NewRequest(
		"PUT",
		test.Addr+"/versions/"+versionID.String(),
		strings.NewReader(`{"name":"new-name"}`),
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
	_version := versionRepository.FindByID(context.TODO(), versionID)
	require.NotNil(t, _version)
	require.Equal(t, newVersionName.String(), _version.Name().String())
}
