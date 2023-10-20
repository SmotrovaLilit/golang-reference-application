package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/optional"
	"strings"
	"testing"
)

//nolint:funlen
func TestDecodeUpdateProgramVersionRequest(t *testing.T) {
	tests := []struct {
		name    string
		request *http.Request
		wantErr error
		want    interface{}
	}{
		{
			name: "valid request",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"new-name", "description": "new-description", "number": "1.0.1"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: nil,
			want: updateprogramversion.NewCommand(
				version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				version.MustNewName("new-name"),
				optional.Of[version.Description](version.MustNewDescription("new-description")),
				optional.Of[version.Number](version.MustNewNumber("1.0.1")),
			),
		},
		{
			name: "valid request with only required fields",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"new-name"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: nil,
			want: updateprogramversion.NewCommand(
				version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				version.MustNewName("new-name"),
				optional.Empty[version.Description](),
				optional.Empty[version.Number](),
			),
		},
		{
			name: "not valid json",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"invalid`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: errInvalidJson,
		},
		{
			name: "invalid version id",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/invalid",
				strings.NewReader(`{"name":"new-name"}`),
			), map[string]string{"id": "invalid"}),
			wantErr: version.ErrInvalidID,
		},
		{
			name: "invalid version name",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"sh"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: version.ErrNameLength,
		},
		{
			name: "invalid version description",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(fmt.Sprintf(`{"name":"new-name", "description": "%s"}`, strings.Repeat("1", 10001))),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: version.ErrDescriptionLength,
		},
		{
			name: "invalid version number",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"new-name", "number": "sh"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: version.ErrInvalidVersionNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeUpdateProgramVersionRequest(slog.Default())(context.TODO(), tt.request)
			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
