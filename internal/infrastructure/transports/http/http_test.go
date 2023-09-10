package http

import (
	"context"
	stdErrors "errors"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"testing"
)

// nolint:funlen
func Test_convertErrorToApiError(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  string
	}{
		{
			name:  "err bad request",
			input: errInvalidJson,
			want:  `{"error":"invalid json","code":"INVALID_JSON"}`,
		},
		{
			name:  "err internal",
			input: stdErrors.New("some error"),
			want:  `{"error":"Internal Server Error","code":"INTERNAL_SERVER_ERROR"}`,
		},
		{
			name:  "err invalid program id",
			input: program.ErrInvalidID,
			want:  `{"error":"invalid program id","code":"INVALID_PROGRAM_ID"}`,
		},
		{
			name:  "err invalid version id",
			input: version.ErrInvalidID,
			want:  `{"error":"invalid version id","code":"INVALID_VERSION_ID"}`,
		},
		{
			name:  "err invalid platform code",
			input: program.ErrInvalidPlatformCode,
			want:  `{"error":"invalid program platform code","code":"INVALID_PROGRAM_PLATFORM_CODE"}`,
		},
		{
			name:  "err update version status",
			input: version.ErrUpdateVersionStatus,
			want:  `{"error":"invalid status to update version","code":"INVALID_STATUS_TO_UPDATE"}`,
		},
		{
			name:  "err version not found",
			input: sharederrors.ErrVersionNotFound,
			want:  `{"error":"version not found","code":"NOT_FOUND"}`,
		},
		{
			name:  "err version name length",
			input: version.ErrNameLength,
			want:  `{"error":"invalid version name length","code":"INVALID_VERSION_NAME_LENGTH"}`,
		},
		{
			name:  "err version description length",
			input: version.ErrDescriptionLength,
			want:  `{"error":"invalid version description length","code":"INVALID_VERSION_DESCRIPTION_LENGTH"}`,
		},
		{
			name:  "err version number is empty",
			input: version.ErrEmptyNumber,
			want:  `{"error":"number is empty","code":"EMPTY_NUMBER"}`,
		},
		{
			name:  "err version description is empty",
			input: version.ErrEmptyDescription,
			want:  `{"error":"description is empty","code":"EMPTY_DESCRIPTION"}`,
		},
		{
			name:  "err invalid status to send to review",
			input: version.ErrInvalidStatusToSendToReview,
			want:  `{"error":"invalid status to send to review","code":"INVALID_STATUS_TO_SEND_TO_REVIEW"}`,
		},
		{
			name:  "err empty number",
			input: version.ErrEmptyNumber,
			want:  `{"error":"number is empty","code":"EMPTY_NUMBER"}`,
		},
		{
			name:  "err invalid status to approve",
			input: version.ErrInvalidStatusToApprove,
			want:  `{"error":"invalid status to approve","code":"INVALID_STATUS_TO_APPROVE"}`,
		},
		{
			name:  "err invalid status to decline",
			input: version.ErrInvalidStatusToDecline,
			want:  `{"error":"invalid status to decline","code":"INVALID_STATUS_TO_DECLINE"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			errorEncoder(context.TODO(), tt.input, writer)
			require.Equal(t, tt.want+"\n", writer.Body.String())
		})
	}
}

func TestErrorEncoderNilError(t *testing.T) {
	require.Panics(t, func() {
		_ = convertErrorToApiError(nil)
	})
}
