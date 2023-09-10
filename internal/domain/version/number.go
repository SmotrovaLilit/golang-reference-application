package version

import (
	"fmt"
	"reference-application/internal/pkg/errorswithcode"
	"regexp"
)

// ErrNumberLength is an error for invalid version number length.
var ErrNumberLength = errorswithcode.NewValidationErrorWithCode("invalid version number length", "INVALID_VERSION_NUMBER_LENGTH")
var ErrEmptyNumber = errorswithcode.NewValidationErrorWithCode("number is empty", "EMPTY_NUMBER")
var ErrInvalidVersionNumber = errorswithcode.NewValidationErrorWithCode("invalid version number", "INVALID_VERSION_NUMBER")

// Number is a type for version number.
type Number string

const (
	numberMaxLength = 15
	numberMinLength = 1
)

var versionRegexp = regexp.MustCompile(`^(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*$`)

// NewNumber is a constructor for Number.
func NewNumber(raw string) (Number, error) {
	if len(raw) < numberMinLength {
		return "", fmt.Errorf(
			"%w: number must be at least %d characters long",
			ErrNumberLength,
			numberMinLength,
		)
	}
	if len(raw) > numberMaxLength {
		return "", fmt.Errorf(
			"%w: number must be at most %d characters long",
			ErrNumberLength,
			numberMaxLength,
		)
	}
	if !versionRegexp.MatchString(raw) {
		return "", ErrInvalidVersionNumber
	}

	return Number(raw), nil
}

// MustNewNumber is a constructor for Number.
// It panics if the given raw string is invalid.
func MustNewNumber(raw string) Number {
	number, err := NewNumber(raw)
	if err != nil {
		panic(err)
	}
	return number
}

// String returns a string representation of a version number.
func (n Number) String() string {
	return string(n)
}
