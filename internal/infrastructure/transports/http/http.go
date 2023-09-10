package http

import (
	"context"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application"
	"reference-application/internal/pkg/errorswithcode"
	xhttp "reference-application/internal/pkg/http"
)

// NewHandler creates a new http.Handler.
func NewHandler(endpoints application.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Handle("/programs", newCreateProgramHandler(endpoints.CreateProgramEndpoint)).Methods(http.MethodPost)
	r.Handle("/versions/{id}", newUpdateProgramVersionHandler(endpoints.UpdateProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/versions/{id}/sendToReview", newSendToReviewProgramVersionHandler(endpoints.SendToReviewProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/versions/{id}/approve", newApproveProgramVersionHandler(endpoints.ApproveProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/versions/{id}/decline", newDeclineProgramVersionHandler(endpoints.DeclineProgramVersionEndpoint)).Methods(http.MethodPut)
	return r
}

// handlersOptions contains the options for the handlers.
var handlersOptions = []kithttp.ServerOption{
	kithttp.ServerErrorEncoder(errorEncoder),
}

// errorEncoder encodes errors.
func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	xhttp.ErrorEncoder(ctx, convertErrorToApiError(err), w)
}

// errInvalidJson is an errors for invalid json.
var errInvalidJson = errorswithcode.New("invalid json", "INVALID_JSON")

// convertErrorToApiError converts an error to an API error with status code.
// If error is unknown, then it returns an internal error.
func convertErrorToApiError(err error) *xhttp.ApiError {
	if err == nil {
		panic("convert error to ApiError: err is nil")
	}

	switch true {
	case errors.As(err, new(*errorswithcode.NotFoundError)):
		return xhttp.NewNotFoundError(err)
	case errors.As(err, new(*errorswithcode.ValidationError)):
		return xhttp.NewValidationError(err)
	case errors.Is(err, errInvalidJson):
		return xhttp.NewBadRequestError(err)
	default:
		// TODO log original error https://github.com/SmotrovaLilit/golang-reference-application/issues/2
		return xhttp.ErrInternal
	}
}
