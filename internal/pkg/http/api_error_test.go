package http

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	errorswithcode "reference-application/internal/pkg/errorswithcode"
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
			name: "error with code",
			args: args{
				err: errorswithcode.New("MESSAGE", "CODE"),
			},
			want: NewApiError(http.StatusUnprocessableEntity, "MESSAGE", "CODE"),
		},
		{
			name: "std error",
			args: args{
				err: errors.New("MESSAGE"),
			},
			want: NewApiError(http.StatusUnprocessableEntity, "MESSAGE", "VALIDATION_ERROR"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewValidationError(tt.args.err)
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
			name: "std error",
			args: args{
				err: errors.New("MESSAGE"),
			},
			want: NewApiError(http.StatusBadRequest, "MESSAGE", "BAD_REQUEST"),
		},
		{
			name: "error with code",
			args: args{
				err: errorswithcode.New("MESSAGE", "CODE"),
			},
			want: NewApiError(http.StatusBadRequest, "MESSAGE", "CODE"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBadRequestError(tt.args.err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewNotFoundError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *ApiError
	}{
		{
			name: "std error",
			args: args{
				err: errors.New("MESSAGE"),
			},
			want: NewApiError(http.StatusNotFound, "MESSAGE", "NOT_FOUND"),
		},
		{
			name: "error with code",
			args: args{
				err: errorswithcode.New("MESSAGE", "CODE"),
			},
			want: NewApiError(http.StatusNotFound, "MESSAGE", "CODE"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotFoundError(tt.args.err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewApiErrorFromError(t *testing.T) {
	type args struct {
		statusCode  int
		err         error
		defaultCode string
	}
	tests := []struct {
		name string
		args args
		want *ApiError
	}{
		{
			name: "error with code",
			args: args{
				statusCode:  http.StatusNotFound,
				err:         errorswithcode.New("MESSAGE", "CODE"),
				defaultCode: "NOT_FOUND",
			},
			want: NewApiError(http.StatusNotFound, "MESSAGE", "CODE"),
		},
		{
			name: "std error",
			args: args{
				statusCode:  http.StatusBadRequest,
				err:         errors.New("MESSAGE"),
				defaultCode: "BAD_REQUEST",
			},
			want: NewApiError(http.StatusBadRequest, "MESSAGE", "BAD_REQUEST"),
		},
		{
			name: "error with code wrapped",
			args: args{
				statusCode:  http.StatusBadRequest,
				err:         fmt.Errorf("%w: new message", errorswithcode.New("ORIGINAL MESSAGE", "CODE")),
				defaultCode: "BAD_REQUEST",
			},
			want: NewApiError(http.StatusBadRequest, "ORIGINAL MESSAGE: new message", "CODE"),
		},
		{
			name: "std error wrapped",
			args: args{
				statusCode:  http.StatusBadRequest,
				err:         fmt.Errorf("%w: new message", errors.New("ORIGINAL MESSAGE")),
				defaultCode: "BAD_REQUEST",
			},
			want: NewApiError(http.StatusBadRequest, "ORIGINAL MESSAGE: new message", "BAD_REQUEST"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApiErrorFromError(tt.args.statusCode, tt.args.err, tt.args.defaultCode)
			require.Equal(t, tt.want, got)
		})
	}
}
