package program

import (
	"reference-application/internal/pkg/id"
)

var (
	// NewID is a constructor for program ID.
	// Example usage:
	//  newID, err := NewID("00000000-0000-0000-0000-000000000000")
	NewID = id.New[*ID]
	// MustNewID is a constructor for program ID.
	// It panics if the given raw string is invalid.
	// Example usage:
	//  newID := MustNew("00000000-0000-0000-0000-000000000000")
	MustNewID = id.MustNew[*ID]
)

// ID is a type for program id.
type ID struct {
	id.Base
}
