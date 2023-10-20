package pager

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLimit(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    Limit
		wantErr error
	}{
		{
			name:  "valid",
			input: 10,
			want:  Limit(10),
		},
		{
			name:    "negative error",
			input:   -1,
			wantErr: ErrNegativeLimit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimit, err := NewLimit(tt.input)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, gotLimit)
		})
	}
}

func TestNewLimitFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Limit
		wantErr error
	}{
		{
			name:  "valid",
			input: "10",
			want:  Limit(10),
		},
		{
			name:  "empty",
			input: "",
			want:  DefaultLimit,
		},
		{
			name:    "parse error",
			input:   "invalid",
			wantErr: ErrParseLimit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimit, err := NewLimitFromString(tt.input)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, gotLimit)
		})
	}
}
