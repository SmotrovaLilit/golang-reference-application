package http

import (
	"context"
	"database/sql"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application"
	"reference-application/internal/pkg/errorswithcode"
	"reference-application/internal/pkg/healthcheck"
	xhttp "reference-application/internal/pkg/http"
	"time"
)

// NewHandler creates a new http.Handler.
func NewHandler(endpoints application.Endpoints, db *sql.DB) http.Handler {
	r := mux.NewRouter()
	healthChecker := healthcheck.New(1*time.Second, 60*time.Second) // TODO make configurable
	healthChecker.Add("database", healthcheck.NewDatabaseChecker(db))
	r.Handle("/health", healthChecker.Handler())
	r.Handle("/ready", healthChecker.Handler())
	r.Handle("/programs", newCreateProgramHandler(endpoints.CreateProgramEndpoint)).Methods(http.MethodPost)
	r.Handle("/versions/{id}", newUpdateProgramVersionHandler(endpoints.UpdateProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/versions/{id}/sendToReview", newSendToReviewProgramVersionHandler(endpoints.SendToReviewProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/versions/{id}/approve", newApproveProgramVersionHandler(endpoints.ApproveProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/versions/{id}/decline", newDeclineProgramVersionHandler(endpoints.DeclineProgramVersionEndpoint)).Methods(http.MethodPut)
	r.Handle("/store/programs", newApprovedProgramsHandler(endpoints.ApprovedProgramsEndpoint)).Methods(http.MethodGet)
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
