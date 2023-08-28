package version

import (
	"reference-application/internal/pkg/id"
)

var (
	// NewID is a constructor for ID.
	// Example usage:
	//  newID, err := NewID("00000000-0000-0000-0000-000000000000")
	NewID = id.New[*ID]
	// MustNewID is a constructor for ID.
	// It panics if the given raw string is invalid.
	// Example usage:
	//  newID := MustNew("00000000-0000-0000-0000-000000000000")
	MustNewID = id.MustNew[*ID]
)

// ID is a type for id.
type ID struct {
	id.Base
}
