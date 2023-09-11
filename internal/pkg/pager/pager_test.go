package pager

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewFromHTTPRequest(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		want    Pager
		wantErr error
	}{
		{
			name: "valid",
			uri:  "/?limit=10&offset=10",
			want: New(
				MustNewLimit(10),
				MustNewOffset(10),
			),
		},
		{
			name: "empty",
			uri:  "/",
			want: Default,
		},
		{
			name:    "parse error",
			uri:     "/?limit=invalid&offset=invalid",
			wantErr: ErrParseLimit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPager, err := NewFromHTTPRequest(httptest.NewRequest(
				http.MethodGet,
				tt.uri,
				nil,
			))
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, gotPager)
		})
	}
}
