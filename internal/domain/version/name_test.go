package version

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		want          Name
		wantErr       error
		wantErrString string
	}{
		{
			name:          "success to create name",
			input:         "smart-calculator",
			want:          "smart-calculator",
			wantErr:       nil,
			wantErrString: "",
		},
		{
			name:          "empty name",
			input:         "",
			want:          "",
			wantErr:       ErrNameLength,
			wantErrString: "invalid version name length: name must be at least 3 characters long",
		},
		{
			name:          "short name",
			input:         "ab",
			want:          "",
			wantErr:       ErrNameLength,
			wantErrString: "invalid version name length: name must be at least 3 characters long",
		},
		{
			name:          "long name",
			input:         strings.Repeat("a", 257),
			want:          "",
			wantErr:       ErrNameLength,
			wantErrString: "invalid version name length: name must be at most 256 characters long",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.input)
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

func TestMustNewName(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want Name
	}{
		{
			name: "success to create name",
			args: args{raw: "smart-calculator"},
			want: "smart-calculator",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MustNewName(tt.args.raw)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMustNewNameFailed(t *testing.T) {
	require.PanicsWithError(t, "invalid version name length: name must be at least 3 characters long", func() {
		_ = MustNewName("")
	})

}
