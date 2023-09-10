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
			raw:     " " + strings.Repeat("1", descriptionMaxLength) + " ",
			want:    MustNewDescription(strings.Repeat("1", descriptionMaxLength)),
			wantErr: nil,
		},
		{
			name:          "empty description",
			raw:           "      ",
			want:          "",
			wantErr:       ErrEmptyDescription,
			wantErrString: "description is empty",
		},
		{
			name:          "long description",
			raw:           strings.Repeat("1", descriptionMaxLength+1),
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

func TestDescription_canSendToReview(t *testing.T) {
	tests := []struct {
		name          string
		description   Description
		wantErr       error
		wantErrString string
	}{
		{
			name:        "success",
			description: MustNewDescription(strings.Repeat("1", descriptionMinLength)),
			wantErr:     nil,
		},
		{
			name:          "short description",
			description:   MustNewDescription(strings.Repeat("1", descriptionMinLength-1)),
			wantErr:       ErrDescriptionLength,
			wantErrString: "invalid version description length: description must be at least 10 characters long",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.description.canSendToReview()
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				require.Equal(t, tt.wantErrString, err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
