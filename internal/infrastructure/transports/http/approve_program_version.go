package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application/commands/approveprogramversion"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
)

// newApproveProgramVersionHandler creates a new http.Handler to send to approve a version.
func newApproveProgramVersionHandler(e approveprogramversion.Endpoint) http.Handler {
	return kithttp.NewServer(
		endpoint.Endpoint(e),
		decodeApproveProgramVersionRequest,
		xhttp.NoContentResponseEncoder,
		getHandlerOptions("ApproveProgramVersion")...,
	)
}

// decodeApproveProgramVersionRequest decodes a request to send to approve a version.
func decodeApproveProgramVersionRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	id, err := version.NewID(mux.Vars(req)["id"])
	if err != nil {
		return nil, err
	}
	return approveprogramversion.NewCommand(id), nil
}
