package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application/commands/declineprogramversion"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
)

// newDeclineProgramVersionHandler creates a new http.Handler to send to decline a version.
func newDeclineProgramVersionHandler(e declineprogramversion.Endpoint) http.Handler {
	return kithttp.NewServer(
		endpoint.Endpoint(e),
		decodeDeclineProgramVersionRequest,
		xhttp.NoContentResponseEncoder,
		handlersOptions...,
	)
}

// decodeDeclineProgramVersionRequest decodes a request to send to decline a version.
func decodeDeclineProgramVersionRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	id, err := version.NewID(mux.Vars(req)["id"])
	if err != nil {
		return nil, err
	}
	return declineprogramversion.NewCommand(id), nil
}
