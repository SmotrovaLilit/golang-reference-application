package version

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewNumber(t *testing.T) {
	tests := []struct {
		name          string
		rawNumber     string
		want          Number
		wantErr       error
		wantErrString string
	}{
		{
			name:          "success to create number",
			rawNumber:     " 1.0.0 ",
			want:          MustNewNumber("1.0.0"),
			wantErr:       nil,
			wantErrString: "",
		},
		{
			name:          "empty number",
			rawNumber:     "   ",
			want:          "",
			wantErr:       ErrNumberLength,
			wantErrString: "invalid version number length: number must be at least 1 characters long",
		},
		{
			name:          "long number",
			rawNumber:     strings.Repeat("1", numberMaxLength+1),
			want:          "",
			wantErr:       ErrNumberLength,
			wantErrString: "invalid version number length: number must be at most 15 characters long",
		},
		{
			name:          "invalid number",
			rawNumber:     " 12-12 ",
			want:          "",
			wantErr:       ErrInvalidVersionNumber,
			wantErrString: "invalid version number",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNumber(tt.rawNumber)
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

func TestMustNewNumberFailed(t *testing.T) {
	require.PanicsWithError(t, "invalid version number", func() {
		_ = MustNewNumber("1--1")
	})

}
