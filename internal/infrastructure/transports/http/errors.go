package http

import "reference-application/internal/pkg/errors"

// ErrInvalidJson is an errors for invalid json.
var ErrInvalidJson = errors.New("invalid json", "INVALID_JSON")
