package http

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorEncoder(t *testing.T) {
	type args struct {
		err error
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
		{
			name:           "success to encode unknown errors",
			args:           args{err: errors.New("message")},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"error":"Internal Server Error","code":"INTERNAL_SERVER_ERROR"}` + "\n",
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
