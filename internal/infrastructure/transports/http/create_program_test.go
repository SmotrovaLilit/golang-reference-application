package http

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/domain/program"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/id"
	"strings"
	"testing"
)

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
				strings.NewReader(`{"id":"3BA2DA12-CF71-49BD-A753-48BE34CD848D","platform_code":"ANDROID"}`),
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
			name: "invalid id",
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
