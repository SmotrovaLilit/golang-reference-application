package http

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/id"
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
			wantErr: xhttp.NewBadRequestError(errors.New("unexpected EOF")),
		},
		{
			name: "invalid version id",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/invalid",
				strings.NewReader(`{"name":"new-name"}`),
			), map[string]string{"id": "invalid"}),
			wantErr: xhttp.NewUnprocessableEntityError(id.ErrInvalidID),
		},
		{
			name: "invalid version name",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411",
				strings.NewReader(`{"name":"sh"}`),
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: xhttp.NewUnprocessableEntityError(version.ErrNameLength),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeUpdateProgramVersionRequest(context.TODO(), tt.request)
			require.Equal(t, tt.wantErr, err)
			// TODO check that the command is created correctly https://github.com/SmotrovaLilit/golang-reference-application/issues/16
		})
	}
}
