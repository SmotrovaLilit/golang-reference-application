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
	if !errors.As(err, &apiErr) {
		err = ErrInternal
	}

	w.WriteHeader(apiErr.StatusCode)
	errEnc := json.NewEncoder(w).Encode(apiErr)
	if errEnc != nil {
		panic(errEnc)
	}
}
