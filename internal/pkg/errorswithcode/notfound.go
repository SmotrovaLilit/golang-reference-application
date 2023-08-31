package errorswithcode

const notFoundDefaultCode = "NOT_FOUND"

type NotFoundError struct {
	err *Error
}

func NewNotFound(message string) *NotFoundError {
	return &NotFoundError{
		err: New(message, notFoundDefaultCode),
	}
}

func NewNotFoundWithCode(message string, code string) *NotFoundError {
	return &NotFoundError{
		err: New(message, code),
	}
}

func (e *NotFoundError) Error() string { return e.err.Error() }

func (e *NotFoundError) Unwrap() error { return e.err }
