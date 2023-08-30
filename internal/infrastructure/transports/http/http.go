package http

import (
	"context"
	stdErrors "errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/id"
)

// NewHandler creates a new http.Handler.
// TODO no one test with this function, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/10
func NewHandler(endpoints application.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Handle("/programs", newCreateProgramHandler(endpoints.CreateProgramEndpoint)).Methods(http.MethodPost)
	r.Handle("/versions/{id}", newUpdateProgramVersionHandler(endpoints.UpdateProgramVersionEndpoint)).Methods(http.MethodPut)
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

// convertErrorToApiError converts an error to an API error with status code.
// If error is unknown, then it returns an internal error.
func convertErrorToApiError(err error) error {
	if err == nil {
		return nil
	}
	switch true {
	// TODO может коды ошибок тоже тут формировать?? чтобы легче было с врапами ошибок
	case stdErrors.Is(err, version.ErrUpdateVersionStatus):
		err = xhttp.NewUnprocessableEntityError(err)
	case stdErrors.Is(err, updateprogramversion.ErrVersionNotFound):
		err = xhttp.NewNotFoundError(err)
	case stdErrors.Is(err, program.ErrInvalidPlatformCode):
		err = xhttp.NewUnprocessableEntityError(err)
	case stdErrors.Is(err, id.ErrInvalidID):
		err = xhttp.NewUnprocessableEntityError(err)
	case stdErrors.Is(err, version.ErrNameLength):
		err = xhttp.NewUnprocessableEntityError(err)
	case stdErrors.Is(err, ErrInvalidJson):
		err = xhttp.NewBadRequestError(err)
	default:
		// TODO log original error
		err = xhttp.ErrInternal
	}
	return err
}
