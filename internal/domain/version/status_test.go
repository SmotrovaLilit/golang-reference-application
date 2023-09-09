package version

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatus_isDraft(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{
			name: "draft status",
			s:    DraftStatus,
			want: true,
		},
		{
			name: "not draft status",
			s:    "NOT_DRAFT",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.IsDraft()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStatus_canUpdate(t *testing.T) {
	tests := []struct {
		name    string
		s       Status
		wantErr error
	}{
		{
			name:    "success",
			s:       DraftStatus,
			wantErr: nil,
		},
		{
			name:    "failed",
			s:       "NOT_DRAFT",
			wantErr: ErrUpdateVersionStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.canUpdate()
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
