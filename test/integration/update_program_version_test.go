package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"reference-application/internal/pkg/optional"
	"reference-application/internal/pkg/tests"
	"strings"
	"testing"
)

func TestUpdateProgramVersionHandler(t *testing.T) {
	// Prepare
	test := tests.PrepareIntegrationTest(t)
	existingVersion := test.PrepareDraftVersion(t)

	newVersionName := version.MustNewName("new-name")
	newVersionDescription := optional.Of[version.Description](version.MustNewDescription("new-description"))

	// Test operation
	req, err := http.NewRequest(
		"PUT",
		test.Addr+"/versions/"+existingVersion.ID().String(),
		strings.NewReader(`{"name":"new-name", "description": "new-description"}`),
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
	versionRepository := repositories.NewVersionRepository(test.DB)
	_version := versionRepository.FindByID(context.TODO(), existingVersion.ID())
	require.NotNil(t, _version)
	require.Equal(t, newVersionName.String(), _version.Name().String())
	require.Equal(t, newVersionDescription.Value().String(), _version.Description().Value().String())
}
