package version

import (
	"fmt"
	"reference-application/internal/pkg/errorswithcode"
)

// ErrNameLength is an error for invalid version name length.
var ErrNameLength = errorswithcode.New("invalid version name length", "INVALID_VERSION_NAME_LENGTH")

// Name is a type for version name.
type Name string

const (
	nameMaxLength = 256
	nameMinLength = 3
)

// NewName is a constructor for Name.
func NewName(raw string) (Name, error) {
	if len(raw) < nameMinLength {
		return "", fmt.Errorf(
			"%w: name must be at least %d characters long",
			ErrNameLength,
			nameMinLength,
		)
	}
	if len(raw) > nameMaxLength {
		return "", fmt.Errorf(
			"%w: name must be at most %d characters long",
			ErrNameLength,
			nameMaxLength,
		)
	}
	return Name(raw), nil
}

// MustNewName is a constructor for Name.
// It panics if the given raw string is invalid.
func MustNewName(raw string) Name {
	name, err := NewName(raw)
	if err != nil {
		panic(err)
	}
	return name
}

// String returns a string representation of a version name.
func (n Name) String() string {
	return string(n)
}
