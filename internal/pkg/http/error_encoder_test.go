package http

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorEncoder(t *testing.T) {
	type args struct {
		err *ApiError
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "success to encode api errors",
			args:           args{err: NewApiError(http.StatusBadRequest, "message", "CODE")},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"error":"message","code":"CODE"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ErrorEncoder(context.TODO(), tt.args.err, w)
			require.Equal(t, tt.wantStatusCode, w.Code)
			require.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}
