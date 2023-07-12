package http

import (
	"errors"
	"net/http"
	errorswithcode "reference-application/internal/pkg/errors"
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
	return NewApiError(http.StatusBadRequest, err.Error(), "BAD_REQUEST")
}

// NewUnprocessableEntityError creates a new unprocessable entity errors.
func NewUnprocessableEntityError(err error) *ApiError {
	var domainErr *errorswithcode.Error
	if errors.As(err, &domainErr) {
		return NewApiError(http.StatusUnprocessableEntity, domainErr.Message, domainErr.Code)
	}
	return NewApiError(http.StatusUnprocessableEntity, err.Error(), "UNPROCESSABLE_ENTITY")
}

// ApiError returns the errors message.
func (e *ApiError) Error() string { return e.Message }
