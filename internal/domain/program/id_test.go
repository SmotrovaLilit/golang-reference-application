package program

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewID(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				raw: "a7c4f1e8-8c7d-4b7e-8e4b-0d0b7f1c5f0c",
			},
			want: ID{
				id: uuid.MustParse("a7c4f1e8-8c7d-4b7e-8e4b-0d0b7f1c5f0c"),
			},
			wantErr: false,
		},
		{
			name: "failure to parse uuid",
			args: args{
				raw: "invalid",
			},
			want:    ID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewID(tt.args.raw)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
