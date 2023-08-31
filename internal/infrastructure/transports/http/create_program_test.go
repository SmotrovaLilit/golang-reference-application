package http

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
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
			wantErr: errInvalidJson,
		},
		{
			name: "invalid program id",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"invalid","platform_code":"ANDROID"}`),
			),
			wantErr: program.ErrInvalidID,
		},
		{
			name: "invalid platform code",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"invalid"}`),
			),
			wantErr: program.ErrInvalidPlatformCode,
		},
		{
			name: "without version",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID"}`),
			),
			wantErr: version.ErrInvalidID,
		},
		{
			name: "invalid version id",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID", "version":{"id":"invalid","name":"smart-calculator"}}`),
			),
			wantErr: version.ErrInvalidID,
		},
		{
			name: "invalid version name",
			request: httptest.NewRequest(
				http.MethodPost,
				"/programs",
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID", "version":{"id":"11A111CF-91F3-49DC-BB6D-AC4235635411","name":"sh"}}`),
			),
			wantErr: version.ErrNameLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeCreateProgramRequest(context.TODO(), tt.request)
			if tt.wantErr == nil {
				require.NoError(t, err)
				return
			}
			require.ErrorIs(t, err, tt.wantErr)
			// TODO check that the command is created correctly https://github.com/SmotrovaLilit/golang-reference-application/issues/16
		})
	}
}
