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

func TestCreateProgramHandler(t *testing.T) {
	test := tests.PrepareIntegrationTest(t)

	req, err := http.NewRequest(
		"POST",
		test.Addr+"/programs",
		strings.NewReader(`{"id":"ecaffa6e-4302-4a46-ae72-44a7bd20dfd5", "platform_code":"ANDROID", "version":{"id":"11a111cf-91f3-49dc-bb6d-ac4235635411","name":"smart-calculator"}}`),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.TODO())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	t.Cleanup(func() {
		resp.Body.Close()
	})

	programRepository := repositories.NewProgramRepository(test.DB)
	versionRepository := repositories.NewVersionRepository(test.DB)
	_program := programRepository.FindByID(context.Background(), program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"))
	require.NotNil(t, _program)
	require.Equal(t, "ecaffa6e-4302-4a46-ae72-44a7bd20dfd5", _program.ID().String())
	require.Equal(t, "ANDROID", _program.PlatformCode().String())
	_version := versionRepository.FindByID(context.Background(), version.MustNewID("11A111CF-91F3-49DC-BB6D-AC4235635411"))
	require.NotNil(t, _version)
	require.Equal(t, "11a111cf-91f3-49dc-bb6d-ac4235635411", _version.ID().String())
	require.Equal(t, "smart-calculator", _version.Name().String())
	require.Equal(t, "ecaffa6e-4302-4a46-ae72-44a7bd20dfd5", _version.ProgramID().String())
	require.Equal(t, "DRAFT", _version.Status().String())
}
