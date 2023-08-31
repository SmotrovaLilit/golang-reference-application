package errorswithcode

const validationDefaultCode = "VALIDATION_ERROR"

// ValidationError is an validation errors with a message and a code.
type ValidationError struct {
	err *Error
}

// NewValidationError creates a new validation errors.
func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		err: New(message, validationDefaultCode),
	}
}

// NewValidationErrorWithCode creates a new validation errors with a code.
func NewValidationErrorWithCode(message string, code string) *ValidationError {
	return &ValidationError{
		err: New(message, code),
	}
}

// Error returns the error message.
func (e *ValidationError) Error() string { return e.err.Error() }

// Unwrap returns the wrapped error.
func (e *ValidationError) Unwrap() error { return e.err }
