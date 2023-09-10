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

func TestStatus_IsOnReview(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{
			name: "success",
			s:    OnReviewStatus,
			want: true,
		},
		{
			name: "failed",
			s:    DraftStatus,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.IsOnReview()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStatus_IsApproved(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{
			name: "success",
			s:    ApprovedStatus,
			want: true,
		},
		{
			name: "failed",
			s:    DraftStatus,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.IsApproved()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStatus_approve(t *testing.T) {
	tests := []struct {
		name    string
		s       Status
		want    Status
		wantErr error
	}{
		{
			name: "success",
			s:    OnReviewStatus,
			want: ApprovedStatus,
		},
		{
			name:    "from_draft",
			s:       DraftStatus,
			wantErr: ErrInvalidStatusToApprove,
		},
		{
			name:    "from_approved",
			s:       ApprovedStatus,
			wantErr: ErrInvalidStatusToApprove,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.approve()
			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStatus_IsDeclined(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{
			name: "success",
			s:    DeclinedStatus,
			want: true,
		},
		{
			name: "failed",
			s:    ApprovedStatus,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.IsDeclined()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStatus_decline(t *testing.T) {
	tests := []struct {
		name    string
		s       Status
		want    Status
		wantErr error
	}{
		{
			name: "success",
			s:    OnReviewStatus,
			want: DeclinedStatus,
		},
		{
			name:    "from_draft",
			s:       DraftStatus,
			wantErr: ErrInvalidStatusToDecline,
		},
		{
			name:    "from_approved",
			s:       ApprovedStatus,
			wantErr: ErrInvalidStatusToDecline,
		},
		{
			name:    "from_declined",
			s:       DeclinedStatus,
			wantErr: ErrInvalidStatusToDecline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.decline()
			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.want, got)
		})
	}
}
func TestMustNewStatus(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want Status
	}{
		{
			name: "draft status",
			raw:  "DRAFT",
			want: DraftStatus,
		},
		{
			name: "on review status",
			raw:  "ON_REVIEW",
			want: OnReviewStatus,
		},
		{
			name: "approved status",
			raw:  "APPROVED",
			want: ApprovedStatus,
		},
		{
			name: "declined status",
			raw:  "DECLINED",
			want: DeclinedStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MustNewStatus(tt.raw)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStatus_Panic(t *testing.T) {
	require.Panics(t, func() {
		_ = MustNewStatus("INVALID")
	})
}
