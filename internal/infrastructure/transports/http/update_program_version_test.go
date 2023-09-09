package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/domain/version"
	"strings"
	"testing"
)

func TestDecodeUpdateProgramVersionRequest(t *testing.T) {
	tests := []struct {
		name    string
		request *http.Request
		wantErr error
	}{
		{
			name: "valid request",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"new-name", "description": "new-description", "number": "1.0.1"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: nil,
		},
		{
			name: "valid request with only required fields",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"new-name"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: nil,
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
				strings.NewReader(`{"name":"new-name", "description": "sh"}`),
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
			_, err := decodeUpdateProgramVersionRequest(context.TODO(), tt.request)
			require.ErrorIs(t, err, tt.wantErr)
			// TODO check that the command is created correctly https://github.com/SmotrovaLilit/golang-reference-application/issues/16
		})
	}
}
