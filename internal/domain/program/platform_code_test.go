package program

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCode(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    PlatformCode
		wantErr bool
	}{
		{
			name:    "success to create android platform code",
			args:    args{raw: "ANDROID"},
			want:    AndroidPlatformCode,
			wantErr: false,
		},
		{
			name:    "success to create iphone platform code",
			args:    args{raw: "IPHONE"},
			want:    IPhonePlatformCode,
			wantErr: false,
		},
		{
			name:    "fail to create platform code",
			args:    args{raw: "invalid"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlatformCode(tt.args.raw)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
