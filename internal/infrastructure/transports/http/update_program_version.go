package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/optional"
)

// newUpdateProgramVersionHandler creates a new http.Handler to update a version.
func newUpdateProgramVersionHandler(e updateprogramversion.Endpoint) http.Handler {
	return kithttp.NewServer(
		endpoint.Endpoint(e),
		decodeUpdateProgramVersionRequest,
		xhttp.NoContentResponseEncoder,
		handlersOptions...,
	)
}

// updateProgramVersionRequestDTO is a DTO for a request to update a version.
type updateProgramVersionRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Number      string `json:"number"`
}

// decodeUpdateProgramVersionRequest decodes a request to update a version.
func decodeUpdateProgramVersionRequest(_ context.Context, req *http.Request) (interface{}, error) {
	id, err := version.NewID(mux.Vars(req)["id"])
	if err != nil {
		return nil, err
	}

	dto := updateProgramVersionRequestDTO{}
	err = json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		// TODO log original error https://github.com/SmotrovaLilit/golang-reference-application/issues/2
		return nil, errInvalidJson
	}

	name, err := version.NewName(dto.Name)
	if err != nil {
		return nil, err
	}

	description := optional.Empty[version.Description]()
	if dto.Description != "" {
		descriptionValue, err := version.NewDescription(dto.Description)
		if err != nil {
			return nil, err
		}
		description = optional.Of[version.Description](descriptionValue)
	}
	number := optional.Empty[version.Number]()
	if dto.Number != "" {
		numberValue, err := version.NewNumber(dto.Number)
		if err != nil {
			return nil, err
		}
		number = optional.Of[version.Number](numberValue)
	}
	return updateprogramversion.NewCommand(id, name, description, number), nil
}
