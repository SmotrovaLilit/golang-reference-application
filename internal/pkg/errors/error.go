package errors

// Error is an errors with a message and a code.
type Error struct {
	Message string
	Code    string
}

// New creates a new errors.
func New(message string, code string) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

// Error returns the errors message.
func (e *Error) Error() string { return e.Message }
