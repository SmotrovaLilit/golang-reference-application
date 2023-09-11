package pager

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewOffset(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    Offset
		wantErr error
	}{
		{
			name:  "valid",
			input: 10,
			want:  Offset(10),
		},
		{
			name:    "negative error",
			input:   -1,
			wantErr: ErrNegativeOffset,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOffset, err := NewOffset(tt.input)
			require.Equal(t, err, tt.wantErr)
			require.Equal(t, tt.want, gotOffset)
		})
	}
}

func TestNewOffsetFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Offset
		wantErr error
	}{
		{
			name:  "valid",
			input: "10",
			want:  Offset(10),
		},
		{
			name:  "empty",
			input: "",
			want:  DefaultOffset,
		},
		{
			name:    "parse error",
			input:   "invalid",
			wantErr: ErrParseOffset,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOffset, err := NewOffsetFromString(tt.input)
			require.Equal(t, err, tt.wantErr)
			require.Equal(t, tt.want, gotOffset)
		})
	}
}
