package pager

import (
	"errors"
	"strconv"
)

// ErrNegativeOffset is an error for negative offset.
var ErrNegativeOffset = errors.New("offset must be positive")

// ErrParseOffset is an error for invalid offset.
var ErrParseOffset = errors.New("offset must be a number")

// DefaultOffset is a default offset when offset is not specified.
var DefaultOffset = MustNewOffset(0)

// Offset is an offset of pager.
type Offset int

// NewOffset is a constructor for Offset.
func NewOffset(raw int) (Offset, error) {
	if raw < 0 {
		return 0, ErrNegativeOffset
	}
	return Offset(raw), nil
}

// MustNewOffset is a constructor for Offset.
// It panics if offset is invalid.
func MustNewOffset(raw int) Offset {
	offset, err := NewOffset(raw)
	if err != nil {
		panic(err)
	}
	return offset
}

// NewOffsetFromString is a constructor for Offset.
func NewOffsetFromString(raw string) (Offset, error) {
	if raw == "" {
		return DefaultOffset, nil
	}
	rawInt, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return 0, ErrParseOffset
	}
	return NewOffset(int(rawInt))
}

// Int returns offset as int.
func (o Offset) Int() int { return int(o) }
