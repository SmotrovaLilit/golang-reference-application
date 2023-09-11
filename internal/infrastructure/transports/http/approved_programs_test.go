package http

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reference-application/internal/application/queries/approvedprograms"
	"reference-application/internal/domain/program"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/pager"
	"testing"
)

func TestDecodeApprovedProgramsRequest(t *testing.T) {
	tt := []struct {
		name    string
		request *http.Request
		wantErr error
		want    approvedprograms.Query
	}{
		{
			name: "valid request without limit and offset",
			request: httptest.NewRequest(
				http.MethodGet,
				"/store/programs",
				nil,
			),
			wantErr: nil,
			want:    approvedprograms.NewQuery(pager.Default),
		},
		{
			name: "valid request with limit and offset",
			request: httptest.NewRequest(
				http.MethodGet,
				"/store/programs?limit=10&offset=20",
				nil,
			),
			wantErr: nil,
			want: approvedprograms.NewQuery(pager.New(
				pager.MustNewLimit(10),
				pager.MustNewOffset(20),
			)),
		},
		{
			name: "invalid pager",
			request: httptest.NewRequest(
				http.MethodGet,
				"/store/programs?limit=invalid",
				nil,
			),
			wantErr: xhttp.NewBadRequestError(pager.ErrParseLimit),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := decodeApprovedProgramsRequest(context.TODO(), tc.request)
			require.Equal(t, tc.wantErr, err)
			if tc.wantErr != nil {
				return
			}
			require.Equal(t, tc.want, got)
		})
	}
}

func TestEncodeApprovedProgramsResponse(t *testing.T) {
	tt := []struct {
		name    string
		request approvedprograms.Result
		wantErr error
		want    string
	}{
		{
			name: "valid request",
			request: []approvedprograms.Program{
				approvedprograms.Program{
					ID:           "e8c95fc8-df71-4e16-9384-9705bf8af74a",
					PlatformCode: program.AndroidPlatformCode.String(),
					Version: approvedprograms.Version{
						ID:          "11a111cf-91f3-49dc-bb6d-ac4235635411",
						Name:        "name",
						Number:      "1.0.0",
						Description: "description",
						Status:      "APPROVED",
					},
				},
				{
					ID:           "93bd8260-9e25-4bb7-8d13-4f4cac60df2a",
					PlatformCode: program.IPhonePlatformCode.String(),
					Version: approvedprograms.Version{
						ID:          "21a111cf-91f3-49dc-bb6d-ac4235635412",
						Name:        "name 1",
						Number:      "1.0.1",
						Description: "description 1",
						Status:      "APPROVED",
					},
				},
			},
			wantErr: nil,
			want:    `{"data":[{"id":"e8c95fc8-df71-4e16-9384-9705bf8af74a","platform_code":"ANDROID","version":[{"id":"11a111cf-91f3-49dc-bb6d-ac4235635411","name":"name","number":"1.0.0","description":"description","status":"APPROVED"}]},{"id":"93bd8260-9e25-4bb7-8d13-4f4cac60df2a","platform_code":"IPHONE","version":[{"id":"21a111cf-91f3-49dc-bb6d-ac4235635412","name":"name 1","number":"1.0.1","description":"description 1","status":"APPROVED"}]}]}` + "\n",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := encodeApprovedProgramsResponse(context.TODO(), w, tc.request)
			require.Equal(t, tc.wantErr, err)
			if tc.wantErr != nil {
				return
			}
			require.Equal(t, tc.want, w.Body.String())
		})
	}
}
