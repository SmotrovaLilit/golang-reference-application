package http

import (
	"context"
	"encoding/json"
	"fmt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/log"
	"reference-application/internal/pkg/optional"
	"reference-application/internal/pkg/resource"
)

// newUpdateProgramVersionHandler creates a new http.Handler to update a version.
func newUpdateProgramVersionHandler(e updateprogramversion.Endpoint, logger *slog.Logger) http.Handler {
	return kithttp.NewServer(
		kitendpoint.Endpoint(e),
		decodeUpdateProgramVersionRequest(logger),
		xhttp.NoContentResponseEncoder,
		getHandlerOptions(e, logger)...,
	)
}

// updateProgramVersionRequestDTO is a DTO for a request to update a version.
type updateProgramVersionRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Number      string `json:"number"`
}

// decodeUpdateProgramVersionRequest decodes a request to update a version.
func decodeUpdateProgramVersionRequest(logger *slog.Logger) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		logger := log.WithContext(ctx, logger)
		id, err := version.NewID(mux.Vars(req)["id"])
		if err != nil {
			logger.Warn(fmt.Sprintf(
				"failed to parse program version id from request params: %s", err.Error()))
			return nil, err
		}
		ctx = resource.PopulateContextWithResourceID(ctx, id.String())
		logger = log.WithContext(ctx, logger)

		dto := updateProgramVersionRequestDTO{}
		err = json.NewDecoder(req.Body).Decode(&dto)
		if err != nil {
			logger.Warn(fmt.Sprintf("invalid request body: %s", err.Error()))
			return nil, errInvalidJson
		}

		name, err := version.NewName(dto.Name)
		if err != nil {
			logger.Warn(err.Error())
			return nil, err
		}

		description := optional.Empty[version.Description]()
		if dto.Description != "" {
			descriptionValue, err := version.NewDescription(dto.Description)
			if err != nil {
				logger.Warn(err.Error())
				return nil, err
			}
			description = optional.Of[version.Description](descriptionValue)
		}
		number := optional.Empty[version.Number]()
		if dto.Number != "" {
			numberValue, err := version.NewNumber(dto.Number)
			if err != nil {
				logger.Warn(err.Error())
				return nil, err
			}
			number = optional.Of[version.Number](numberValue)
		}
		return updateprogramversion.NewCommand(id, name, description, number), nil
	}
}
