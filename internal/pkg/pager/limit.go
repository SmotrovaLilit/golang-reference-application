package pager

import (
	"errors"
	"strconv"
)

// ErrNegativeLimit is an error for negative limit.
var ErrNegativeLimit = errors.New("limit must be positive")

// ErrParseLimit is an error for invalid limit.
var ErrParseLimit = errors.New("limit must be a number")

// DefaultLimit is a default limit when limit is not specified.
var DefaultLimit = MustNewLimit(10)

// Limit is a limit of pager.
type Limit int

// NewLimit is a constructor for Limit.
func NewLimit(raw int) (Limit, error) {
	if raw < 0 {
		return 0, ErrNegativeLimit
	}
	return Limit(raw), nil
}

// MustNewLimit is a constructor for Limit.
// It panics if limit is invalid.
func MustNewLimit(raw int) Limit {
	limit, err := NewLimit(raw)
	if err != nil {
		panic(err)
	}
	return limit
}

// NewLimitFromString is a constructor for Limit.
// It parses limit from string.
// If limit is empty, it returns default limit.
func NewLimitFromString(raw string) (Limit, error) {
	if raw == "" {
		return DefaultLimit, nil
	}
	rawInt, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return 0, ErrParseLimit
	}
	return NewLimit(int(rawInt))
}

// Int returns limit as int.
func (l Limit) Int() int { return int(l) }
