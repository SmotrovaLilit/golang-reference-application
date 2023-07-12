package http

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoContentResponseEncoder(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success to encode no content response",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			err := NoContentResponseEncoder(context.TODO(), writer, nil)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, http.StatusNoContent, writer.Code)
			require.Equal(t, "", writer.Body.String())
		})
	}
}
