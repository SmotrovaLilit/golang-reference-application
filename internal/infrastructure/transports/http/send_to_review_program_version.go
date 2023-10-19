package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application/commands/sendtoreviewprogramversion"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
)

// newSendToReviewProgramVersionHandler creates a new http.Handler to send to review a version.
func newSendToReviewProgramVersionHandler(e sendtoreviewprogramversion.Endpoint) http.Handler {
	return kithttp.NewServer(
		endpoint.Endpoint(e),
		decodeSendToReviewProgramVersionRequest,
		xhttp.NoContentResponseEncoder,
		getHandlerOptions("SendToReviewProgramVersion")...,
	)
}

// decodeSendToReviewProgramVersionRequest decodes a request to send to review a version.
func decodeSendToReviewProgramVersionRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	id, err := version.NewID(mux.Vars(req)["id"])
	if err != nil {
		return nil, err
	}
	return sendtoreviewprogramversion.NewCommand(id), nil
}
