package http

import (
	"context"
	"encoding/json"
	"net/http"
)

// ErrorEncoder encodes an errors to the response.
func ErrorEncoder(_ context.Context, err *ApiError, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(err.StatusCode)
	errEnc := json.NewEncoder(w).Encode(err)
	if errEnc != nil {
		panic(errEnc)
	}
}
