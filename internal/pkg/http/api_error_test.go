package http

import (
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	errorswithcode "reference-application/internal/pkg/errors"
	"testing"
)

func TestNewUnprocessableEntityError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *ApiError
	}{
		{
			name: "domain error",
			args: args{
				err: errorswithcode.New("MESSAGE", "CODE"),
			},
			want: NewApiError(http.StatusUnprocessableEntity, "MESSAGE", "CODE"),
		},
		{
			name: "general error",
			args: args{
				err: errors.New("MESSAGE"),
			},
			want: NewApiError(http.StatusUnprocessableEntity, "MESSAGE", "UNPROCESSABLE_ENTITY"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnprocessableEntityError(tt.args.err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewBadRequestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *ApiError
	}{
		{
			name: "case1",
			args: args{
				err: errors.New("MESSAGE"),
			},
			want: NewApiError(http.StatusBadRequest, "MESSAGE", "BAD_REQUEST"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBadRequestError(tt.args.err)
			require.Equal(t, tt.want, got)
		})
	}
}
