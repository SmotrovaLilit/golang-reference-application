package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// ErrorEncoder encodes an errors to the response.
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var apiErr *ApiError
	if errors.As(err, &apiErr) {
		w.WriteHeader(apiErr.StatusCode)
		errEnc := json.NewEncoder(w).Encode(apiErr)
		if errEnc != nil {
			panic(errEnc)
		}
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	errEnc := json.NewEncoder(w).Encode(ApiError{
		Message: http.StatusText(http.StatusInternalServerError),
		Code:    "INTERNAL_SERVER_ERROR",
	})
	if errEnc != nil {
		panic(errEnc)
	}
}
