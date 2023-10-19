package http

import (
	"context"
	"fmt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"reference-application/internal/application/commands/declineprogramversion"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/log"
)

// newDeclineProgramVersionHandler creates a new http.Handler to send to decline a version.
func newDeclineProgramVersionHandler(e declineprogramversion.Endpoint, logger *slog.Logger) http.Handler {
	return kithttp.NewServer(
		kitendpoint.Endpoint(e),
		decodeDeclineProgramVersionRequest(logger),
		xhttp.NoContentResponseEncoder,
		getHandlerOptions(e, logger)...,
	)
}

// decodeDeclineProgramVersionRequest decodes a request to send to decline a version.
func decodeDeclineProgramVersionRequest(logger *slog.Logger) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		logger := log.WithContext(ctx, logger)
		id, err := version.NewID(mux.Vars(req)["id"])
		if err != nil {
			logger.Warn(fmt.Sprintf(
				"failed to parse program version id from request params: %s", err.Error()))
			return nil, err
		}
		return declineprogramversion.NewCommand(id), nil
	}
}
