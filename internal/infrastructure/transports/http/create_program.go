package http

import (
	"context"
	"encoding/json"
	"fmt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"log/slog"
	"net/http"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/log"
	"reference-application/internal/pkg/resource"
)

// newCreateProgramHandler creates a new http.Handler to create a program.
func newCreateProgramHandler(e createprogram.Endpoint, logger *slog.Logger) http.Handler {
	return kithttp.NewServer(
		kitendpoint.Endpoint(e),
		decodeCreateProgramRequest(logger),
		xhttp.NoContentResponseEncoder,
		getHandlerOptions(e, logger)...,
	)
}

// createProgramRequestDTO is a DTO for a request to create a program.
type createProgramRequestDTO struct {
	ID           string `json:"id"`
	PlatformCode string `json:"platform_code"`
	Version      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"version"`
}

// decodeCreateProgramRequest decodes a request to create a program.
func decodeCreateProgramRequest(logger *slog.Logger) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, request *http.Request) (interface{}, error) {
		logger := log.WithContext(ctx, logger)
		dto := createProgramRequestDTO{}
		err := json.NewDecoder(request.Body).Decode(&dto)
		if err != nil {
			logger.Warn(fmt.Sprintf("invalid request body: %s", err.Error()))
			return nil, errInvalidJson
		}
		_id, err := program.NewID(dto.ID)
		if err != nil {
			logger.Warn(err.Error())
			return nil, err
		}
		ctx = resource.PopulateContextWithResourceID(ctx, _id.String())
		logger = log.WithContext(ctx, logger)
		platformCode, err := program.NewPlatformCode(dto.PlatformCode)
		if err != nil {
			logger.Warn(err.Error())
			return nil, err
		}
		versionID, err := version.NewID(dto.Version.ID)
		if err != nil {
			logger.Warn(err.Error())
			return nil, err
		}
		versionName, err := version.NewName(dto.Version.Name)
		if err != nil {
			logger.Warn(err.Error())
			return nil, err
		}

		return createprogram.NewCommand(
			_id,
			platformCode,
			versionID,
			versionName,
		), nil
	}
}
