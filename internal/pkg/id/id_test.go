package id

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reference-application/internal/pkg/errorswithcode"
	"testing"
)

type testID struct {
	Base
}

var testErrInvalidID = errorswithcode.New("invalid test id", "INVALID_TEST_ID")
var newTestID, mustNewTestID = ConstructorsFactory[*testID](testErrInvalidID)

func TestParse(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    testID
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				raw: "a7c4f1e8-8c7d-4b7e-8e4b-0d0b7f1c5f0c",
			},
			want:    testID{Base: Base{uuid: uuid.MustParse("a7c4f1e8-8c7d-4b7e-8e4b-0d0b7f1c5f0c")}},
			wantErr: false,
		},
		{
			name: "failure to parse id",
			args: args{
				raw: "invalid",
			},
			want:    testID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTestID(tt.args.raw)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestID_String(t *testing.T) {
	tests := []struct {
		name string
		i    testID
		want string
	}{
		{
			name: "success",
			i:    mustNewTestID("a7c4f1e8-8c7d-4b7e-8e4b-0d0b7f1c5f0c"),
			want: "a7c4f1e8-8c7d-4b7e-8e4b-0d0b7f1c5f0c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.String()
			require.Equal(t, tt.want, got)
		})
	}
}
