package http

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/id"
	"strings"
	"testing"
)

//nolint:funlen
func TestDecodeCreateProgramRequest(t *testing.T) {
	tests := []struct {
		name    string
		request *http.Request
		wantErr error
	}{
		{
			name: "valid request",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID", "version":{"id":"11A111CF-91F3-49DC-BB6D-AC4235635411","name":"smart-calculator"}}`),
			),
		},
		{
			name: "not valid json",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"invalid`),
			),
			wantErr: xhttp.NewBadRequestError(errors.New("unexpected EOF")),
		},
		{
			name: "invalid program id",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"invalid","platform_code":"ANDROID"}`),
			),
			wantErr: xhttp.NewUnprocessableEntityError(id.ErrInvalidID),
		},
		{
			name: "invalid platform code",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"invalid"}`),
			),
			wantErr: xhttp.NewUnprocessableEntityError(program.ErrInvalidPlatformCode),
		},
		{
			name: "without version",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID"}`),
			),
			wantErr: xhttp.NewUnprocessableEntityError(id.ErrInvalidID),
		},
		{
			name: "invalid version id",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID", "version":{"id":"invalid","name":"smart-calculator"}}`),
			),
			wantErr: xhttp.NewUnprocessableEntityError(id.ErrInvalidID),
		},
		{
			name: "invalid version name",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID", "version":{"id":"11A111CF-91F3-49DC-BB6D-AC4235635411","name":"sh"}}`),
			),
			wantErr: xhttp.NewUnprocessableEntityError(version.ErrNameLength),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeCreateProgramRequest(context.TODO(), tt.request)
			if tt.wantErr == nil {
				require.NoError(t, err)
				return
			}
			require.Equal(t, tt.wantErr, err)
		})
	}
}
