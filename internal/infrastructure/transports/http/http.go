package http

import (
	"context"
	"database/sql"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"reference-application/internal/application"
	"reference-application/internal/pkg/errorswithcode"
	"reference-application/internal/pkg/healthcheck"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/log"
	"reference-application/internal/pkg/resource"
	"time"
)

// NewHandler creates a new http.Handler.
func NewHandler(endpoints application.Endpoints, db *sql.DB, l *slog.Logger) http.Handler {
	router := mux.NewRouter()
	healthChecker := healthcheck.New(1*time.Second, 60*time.Second) // TODO make configurable
	healthChecker.Add("database", healthcheck.NewDatabaseChecker(db))
	router.Handle("/health", healthChecker.Handler())
	router.Handle("/ready", healthChecker.Handler())
	router.Handle("/programs", newCreateProgramHandler(endpoints.CreateProgramEndpoint, l)).Methods(http.MethodPost)
	router.Handle("/versions/{id}", newUpdateProgramVersionHandler(endpoints.UpdateProgramVersionEndpoint, l)).Methods(http.MethodPut)
	router.Handle("/versions/{id}/sendToReview", newSendToReviewProgramVersionHandler(endpoints.SendToReviewProgramVersionEndpoint, l)).Methods(http.MethodPut)
	router.Handle("/versions/{id}/approve", newApproveProgramVersionHandler(endpoints.ApproveProgramVersionEndpoint, l)).Methods(http.MethodPut)
	router.Handle("/versions/{id}/decline", newDeclineProgramVersionHandler(endpoints.DeclineProgramVersionEndpoint, l)).Methods(http.MethodPut)
	router.Handle("/store/programs", newApprovedProgramsHandler(endpoints.ApprovedProgramsEndpoint, l)).Methods(http.MethodGet)

	return router
}

type endpointType interface {
	~func(ctx context.Context, request interface{}) (response interface{}, err error)
}

func getHandlerOptions[T endpointType](e T, logger *slog.Logger) []kithttp.ServerOption {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(errorEncoder(logger)),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	}
	if re, ok := any(e).(resource.Endpoint); ok {
		opts = append(opts, kithttp.ServerBefore(resource.PopulateContextRequestFunc(
			re.ResourceName(),
			re.ResourceAction(),
		)))
	}
	return opts
}

// errorEncoder encodes errors.

func errorEncoder(logger *slog.Logger) func(ctx context.Context, err error, w http.ResponseWriter) {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		xhttp.ErrorEncoder(ctx, convertErrorToApiError(ctx, logger, err), w)
	}
}

// errInvalidJson is an errors for invalid json.
var errInvalidJson = errorswithcode.New("invalid json", "INVALID_JSON")

// convertErrorToApiError converts an error to an API error with status code.
// If error is unknown, then it returns an internal error.
func convertErrorToApiError(ctx context.Context, logger *slog.Logger, err error) *xhttp.ApiError {
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
		logger := log.WithContext(ctx, logger)
		logger.Error(err.Error())
		return xhttp.ErrInternal
	}
}
