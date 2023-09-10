package version

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewDescription(t *testing.T) {
	tests := []struct {
		name          string
		raw           string
		want          Description
		wantErr       error
		wantErrString string
	}{
		{
			name:    "success to create description",
			raw:     "sun",
			want:    MustNewDescription("sun"),
			wantErr: nil,
		},
		{
			name:          "empty description",
			raw:           "",
			want:          "",
			wantErr:       ErrEmptyDescription,
			wantErrString: "description is empty",
		},
		{
			name:          "long description",
			raw:           strings.Repeat("1", 10001),
			want:          "",
			wantErr:       ErrDescriptionLength,
			wantErrString: "invalid version description length: description must be at most 10000 characters long",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDescription(tt.raw)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				require.Equal(t, tt.wantErrString, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
