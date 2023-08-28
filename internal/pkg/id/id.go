package id

import (
	"github.com/google/uuid"
	"reference-application/internal/pkg/errors"
)

// ErrInvalidID is an errors for invalid id.
var ErrInvalidID = errors.New("invalid id", "INVALID_ID")

// Base is a base type for id.
type Base struct {
	uuid uuid.UUID
}

func (b *Base) setUUID(_uuid uuid.UUID) { b.uuid = _uuid }

// String returns a string representation of ID.
func (b Base) String() string { return b.uuid.String() }

type id[T any] interface {
	*T
	setUUID(uid uuid.UUID)
}

// New is a constructor for ID.
func New[PT id[T], T any](raw string) (T, error) {
	var newID T
	var ptrNewID = PT(&newID)

	uid, err := uuid.Parse(raw)
	if err != nil {
		return newID, ErrInvalidID
	}
	ptrNewID.setUUID(uid)

	return newID, nil
}

// MustNew is a constructor for ID.
// It panics if the given raw string is invalid.
func MustNew[PT id[T], T any](raw string) T {
	newID, err := New[PT, T](raw)
	if err != nil {
		panic(err)
	}
	return newID
}
