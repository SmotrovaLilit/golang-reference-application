package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/domain/version"
	"testing"
)

func Test_decodeSendToReviewProgramVersionRequest(t *testing.T) {
	tests := []struct {
		name    string
		request *http.Request
		wantErr error
	}{
		{
			name: "valid request",
			request: mux.SetURLVars(httptest.NewRequest(
				http.MethodPut,
				"/versions/11a111cf-91f3-49dc-bb6d-ac4235635411/sendToReview",
				nil,
			), map[string]string{"id": "11a111cf-91f3-49dc-bb6d-ac4235635411"}),
			wantErr: nil,
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
				_, err := decodeSendToReviewProgramVersionRequest(context.TODO(), tt.request)
				require.ErrorIs(t, err, tt.wantErr)
				// TODO check that the command is created correctly https://github.com/SmotrovaLilit/golang-reference-application/issues/16
			})
		})
	}
}
