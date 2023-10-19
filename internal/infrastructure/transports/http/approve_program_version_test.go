package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/application/commands/approveprogramversion"
	"reference-application/internal/domain/version"
	"testing"
)

func Test_decodeApproveProgramVersionRequest(t *testing.T) {
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
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411/sendToReview",
				nil,
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: nil,
			want: approveprogramversion.NewCommand(
				version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
			),
		},
		{
			name: "invalid version id",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411/sendToReview",
				nil,
			), map[string]string{"id": "invalid id"}),
			wantErr: version.ErrInvalidID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				got, err := decodeApproveProgramVersionRequest(slog.Default())(context.TODO(), tt.request)
				if tt.wantErr == nil {
					require.NoError(t, err)
					require.Equal(t, tt.want, got)
				} else {
					require.ErrorIs(t, err, tt.wantErr)
				}
			})
		})
	}
}
