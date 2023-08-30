package http

import (
	stdErrors "errors"
	"github.com/stretchr/testify/require"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/id"
	"testing"
)

func Test_convertErrorToApiError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			wantErr: nil,
		},
		{
			name: "err bad request",
			args: args{
				err: ErrInvalidJson,
			},
			wantErr: xhttp.NewBadRequestError(ErrInvalidJson),
		},
		{
			name: "err internal",
			args: args{
				err: stdErrors.New("some error"),
			},
			wantErr: xhttp.ErrInternal,
		},
		{
			name: "err invalid id",
			args: args{
				err: id.ErrInvalidID,
			},
			wantErr: xhttp.NewUnprocessableEntityError(id.ErrInvalidID),
		},
		{
			name: "err invalid platform code",
			args: args{
				err: program.ErrInvalidPlatformCode,
			},
			wantErr: xhttp.NewUnprocessableEntityError(program.ErrInvalidPlatformCode),
		},
		{
			name: "err update version status",
			args: args{
				err: version.ErrUpdateVersionStatus,
			},
			wantErr: xhttp.NewUnprocessableEntityError(version.ErrUpdateVersionStatus),
		},
		{
			name: "err version not found",
			args: args{
				err: updateprogramversion.ErrVersionNotFound,
			},
			wantErr: xhttp.NewNotFoundError(updateprogramversion.ErrVersionNotFound),
		},
		{
			name: "err version name length",
			args: args{
				err: version.ErrNameLength,
			},
			wantErr: xhttp.NewUnprocessableEntityError(version.ErrNameLength),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := convertErrorToApiError(tt.args.err)
			require.Equal(t, tt.wantErr, err)
		})
	}
}
