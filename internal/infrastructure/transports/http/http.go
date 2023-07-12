package http

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"reference-application/internal/application"
	xhttp "reference-application/internal/pkg/http"
)

// NewHandler creates a new http.Handler.
func NewHandler(endpoints application.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Handle("/programs", newCreateProgramHandler(endpoints.CreateProgramEndpoint)).Methods(http.MethodPost)
	return r
}

// handlersOptions contains the options for the handlers.
var handlersOptions = []kithttp.ServerOption{
	kithttp.ServerErrorEncoder(xhttp.ErrorEncoder),
}
