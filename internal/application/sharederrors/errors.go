package sharederrors

import "reference-application/internal/pkg/errorswithcode"

var (
	// ErrVersionNotFound is a version not found errors.
	ErrVersionNotFound = errorswithcode.NewNotFound("version not found")
)
