package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
)

// newCreateProgramHandler creates a new http.Handler to create a program.
// TODO no one test with this function
func newCreateProgramHandler(e createprogram.Endpoint) http.Handler {
	return kithttp.NewServer(
		endpoint.Endpoint(e),
		decodeCreateProgramRequest,
		xhttp.NoContentResponseEncoder,
		handlersOptions...,
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
func decodeCreateProgramRequest(_ context.Context, request *http.Request) (interface{}, error) {
	dto := createProgramRequestDTO{}
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		return nil, xhttp.NewBadRequestError(err)
	}
	_id, err := program.NewID(dto.ID)
	if err != nil {
		return nil, xhttp.NewUnprocessableEntityError(err)
	}
	platformCode, err := program.NewPlatformCode(dto.PlatformCode)
	if err != nil {
		return nil, xhttp.NewUnprocessableEntityError(err)
	}
	versionID, err := version.NewID(dto.Version.ID)
	if err != nil {
		return nil, xhttp.NewUnprocessableEntityError(err)
	}
	versionName, err := version.NewName(dto.Version.Name)
	if err != nil {
		return nil, xhttp.NewUnprocessableEntityError(err)
	}

	return createprogram.NewCommand(
		_id,
		platformCode,
		versionID,
		versionName,
	), nil
}
