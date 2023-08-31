package id

import (
	"github.com/google/uuid"
)

// Base is a base type for id.
type Base struct {
	uuid uuid.UUID
}

func (b *Base) setUUID(_uuid uuid.UUID) { b.uuid = _uuid }

// String returns a string representation of ID.
func (b Base) String() string { return b.uuid.String() }

type idSetter[T any] interface {
	*T
	setUUID(uid uuid.UUID)
}

type NewFunc[T any] func(raw string) (T, error)

func NewFactory[PT idSetter[T], T any](errInvalidID error) NewFunc[T] {
	return func(raw string) (T, error) {
		var newID T
		var ptrNewID = PT(&newID)

		uid, err := uuid.Parse(raw)
		if err != nil {
			return newID, errInvalidID
		}
		ptrNewID.setUUID(uid)

		return newID, nil
	}
}

type MustNewFunc[T any] func(raw string) T

func MustNewFactory[PT idSetter[T], T any](errInvalidID error) MustNewFunc[T] {
	return func(raw string) T {
		newID, err := NewFactory[PT, T](errInvalidID)(raw)
		if err != nil {
			panic(err)
		}
		return newID
	}
}

// ConstructorsFactory is a generic function for creating constructors for ID.
// It returns New and MustNew functions.
func ConstructorsFactory[PT idSetter[T], T any](errInvalidID error) (
	NewFunc[T],
	MustNewFunc[T],
) {
	return NewFactory[PT, T](errInvalidID),
		MustNewFactory[PT, T](errInvalidID)
}
