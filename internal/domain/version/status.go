package version

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
