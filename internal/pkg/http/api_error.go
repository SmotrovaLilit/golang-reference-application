package http

import (
	"errors"
	"net/http"
	errorswithcode "reference-application/internal/pkg/errors"
)

var (
	// ErrInternal is an internal server errors.
	ErrInternal = NewApiError(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"INTERNAL_SERVER_ERROR",
	)
)

// ApiError is an errors with a status code, a message and a code.
type ApiError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error,omitempty"`
	Code       string `json:"code,omitempty"`
}

// NewApiError creates a new errors.
func NewApiError(statusCode int, message string, code string) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		Message:    message,
		Code:       code,
	}
}

// NewBadRequestError creates a new bad request errors.
func NewBadRequestError(err error) *ApiError {
	var errorWithCode *errorswithcode.Error
	if errors.As(err, &errorWithCode) {
		return NewApiError(http.StatusBadRequest, errorWithCode.Message, errorWithCode.Code)
	}
	return NewApiError(http.StatusBadRequest, err.Error(), "BAD_REQUEST")
}

// NewUnprocessableEntityError creates a new unprocessable entity errors.
func NewUnprocessableEntityError(err error) *ApiError {
	var errorWithCode *errorswithcode.Error
	if errors.As(err, &errorWithCode) {
		return NewApiError(http.StatusUnprocessableEntity, errorWithCode.Message, errorWithCode.Code)
	}
	return NewApiError(http.StatusUnprocessableEntity, err.Error(), "UNPROCESSABLE_ENTITY")
}

// NewNotFoundError creates a new not found errors.
func NewNotFoundError(err error) *ApiError {
	var errorWithCode *errorswithcode.Error
	if errors.As(err, &errorWithCode) {
		return NewApiError(http.StatusNotFound, errorWithCode.Message, errorWithCode.Code)
	}
	return NewApiError(http.StatusNotFound, err.Error(), "NOT_FOUND")
}

// ApiError returns the errors message.
func (e *ApiError) Error() string { return e.Message }
