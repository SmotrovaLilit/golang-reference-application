package program

import (
	"github.com/google/uuid"
	"reference-application/internal/pkg/errors"
)

// ErrInvalidID is an errors for invalid program id.
var ErrInvalidID = errors.New("invalid program id", "INVALID_PROGRAM_ID")

// ID is a type for program id.
type ID struct {
	id uuid.UUID
}

// NewID is a constructor for ID.
func NewID(raw string) (ID, error) {
	id, err := uuid.Parse(raw)
	if err != nil {
		return ID{}, ErrInvalidID
	}
	return ID{id: id}, nil
}

// String returns a string representation of ID.
func (i ID) String() string { return i.id.String() }
