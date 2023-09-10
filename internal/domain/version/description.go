package version

import (
	"fmt"
	"reference-application/internal/pkg/errorswithcode"
)

// ErrDescriptionLength is an error for invalid version description length.
var ErrDescriptionLength = errorswithcode.NewValidationErrorWithCode("invalid version description length", "INVALID_VERSION_DESCRIPTION_LENGTH")
var ErrEmptyDescription = errorswithcode.NewValidationErrorWithCode("description is empty", "EMPTY_DESCRIPTION")

// Description is a type for version description.
type Description string

const (
	descriptionMaxLength = 10000
	descriptionMinLength = 10
)

// NewDescription is a constructor for Description.
func NewDescription(raw string) (Description, error) {
	if raw == "" {
		return "", ErrEmptyDescription
	}
	if len(raw) > descriptionMaxLength {
		return "", fmt.Errorf(
			"%w: description must be at most %d characters long",
			ErrDescriptionLength,
			descriptionMaxLength,
		)
	}
	return Description(raw), nil
}

// MustNewDescription is a constructor for Description.
// It panics if the given raw string is invalid.
func MustNewDescription(raw string) Description {
	description, err := NewDescription(raw)
	if err != nil {
		panic(err)
	}
	return description
}

// String returns a string representation of a version description.
func (n Description) String() string {
	return string(n)
}

// canSendToReview checks if a version description allows to send version to review.
func (n Description) canSendToReview() error {
	if len(n.String()) < descriptionMinLength {
		return fmt.Errorf(
			"%w: description must be at least %d characters long",
			ErrDescriptionLength,
			descriptionMinLength,
		)
	}
	return nil
}
