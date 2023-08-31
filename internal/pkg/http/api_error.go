package http

import (
	"errors"
	"net/http"
	"reference-application/internal/pkg/errorswithcode"
)

const (
	badRequestCode          = "BAD_REQUEST"
	unprocessableEntityCode = "UNPROCESSABLE_ENTITY"
	notFoundCode            = "NOT_FOUND"
	internalServerError     = "INTERNAL_SERVER_ERROR"
)

var (
	// ErrInternal is an internal server errors.
	ErrInternal = NewApiError(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		internalServerError,
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

// NewApiErrorFromError creates a new API errors from an error.
func NewApiErrorFromError(statusCode int, err error, defaultCode string) *ApiError {
	var errorWithCode *errorswithcode.Error
	if errors.As(err, &errorWithCode) {
		return NewApiError(statusCode, err.Error(), errorWithCode.Code)
	}
	return NewApiError(statusCode, err.Error(), defaultCode)
}

// NewBadRequestError creates a new bad request errors.
func NewBadRequestError(err error) *ApiError {
	return NewApiErrorFromError(http.StatusBadRequest, err, badRequestCode)
}

// NewUnprocessableEntityError creates a new unprocessable entity errors.
func NewUnprocessableEntityError(err error) *ApiError {
	return NewApiErrorFromError(http.StatusUnprocessableEntity, err, unprocessableEntityCode)
}

// NewNotFoundError creates a new not found errors.
func NewNotFoundError(err error) *ApiError {
	return NewApiErrorFromError(http.StatusNotFound, err, notFoundCode)
}

// ApiError returns the errors message.
func (e *ApiError) Error() string { return e.Message }
