package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"reference-application/internal/domain/program"
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
		strings.NewReader(`{"id":"ecaffa6e-4302-4a46-ae72-44a7bd20dfd5", "platform_code":"ANDROID"}`),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.TODO())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	programRepository := repositories.NewProgramRepository(test.DB)
	_program := programRepository.FindByID(context.Background(), program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"))
	require.NotNil(t, _program)
	require.Equal(t, "ecaffa6e-4302-4a46-ae72-44a7bd20dfd5", _program.ID().String())
	require.Equal(t, "ANDROID", _program.PlatformCode().String())
}
