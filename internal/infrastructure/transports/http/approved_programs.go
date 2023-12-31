package http

import (
	"context"
	"fmt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"log/slog"
	"net/http"
	"reference-application/internal/application/queries/approvedprograms"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/log"
	"reference-application/internal/pkg/pager"
	"reference-application/internal/pkg/slices"
)

// programsHandlerOptions is a list of options for approved programs handler.
func newApprovedProgramsHandler(e approvedprograms.Endpoint, logger *slog.Logger) http.Handler {
	return kithttp.NewServer(
		kitendpoint.Endpoint(e),
		decodeApprovedProgramsRequest(logger),
		encodeApprovedProgramsResponse,
		getHandlerOptions(e, logger)...,
	)
}

// decodeApprovedProgramsRequest decodes request from endpoint.
func decodeApprovedProgramsRequest(logger *slog.Logger) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		logger := log.WithContext(ctx, logger)
		_pager, err := pager.NewFromHTTPRequest(r)
		if err != nil {
			logger.Warn(fmt.Sprintf("parsing of query parameters is failed: %s", err.Error()))
			return nil, xhttp.NewBadRequestError(err)
		}
		return approvedprograms.NewQuery(_pager), nil
	}
}

// approvedProgramVersionDTO is a DTO for a program version.
type approvedProgramVersionDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// approvedProgramDTO is a DTO for a program.
type approvedProgramDTO struct {
	ID           string                      `json:"id"`
	PlatformCode string                      `json:"platform_code"`
	Version      []approvedProgramVersionDTO `json:"version"`
}

// approvedProgramsResponseDTO is a DTO for programs response.
type approvedProgramsResponseDTO struct {
	Data []approvedProgramDTO `json:"data"`
}

// encodeApprovedProgramsResponse encodes response from endpoint.
func encodeApprovedProgramsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	queryResult := response.(approvedprograms.Result)
	dto := approvedProgramsResponseDTO{
		Data: slices.Convert(func(program approvedprograms.Program) approvedProgramDTO {
			return approvedProgramDTO{
				ID:           program.ID,
				PlatformCode: program.PlatformCode,
				Version: []approvedProgramVersionDTO{{
					ID:          program.Version.ID,
					Name:        program.Version.Name,
					Number:      program.Version.Number,
					Description: program.Version.Description,
					Status:      program.Version.Status,
				}},
			}
		}, queryResult),
	}
	return kithttp.EncodeJSONResponse(ctx, w, dto)
}
