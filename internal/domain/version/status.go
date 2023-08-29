package version

import "reference-application/internal/pkg/errors"

var (
	ErrUpdateStatus = errors.New("invalid status to update version", "INVALID_STATUS_TO_UPDATE")
)

// Status is a type for a version status.
type Status string

const (
	DraftStatus Status = "DRAFT"
)

// String returns a string representation of a status.
func (s Status) String() string {
	return string(s)
}

// MustNewStatus create a status from a string.
func MustNewStatus(s string) Status {
	switch s {
	case DraftStatus.String():
		return DraftStatus
	}
	panic("unknown status")
}

// isDraft checks if a version status is draft.
func (s Status) isDraft() bool {
	return s == DraftStatus
}

// canUpdate checks if a version status allows to update version.
func (s Status) canUpdate() error {
	if s.isDraft() {
		return nil
	}
	return ErrUpdateStatus
}
