package version

import "reference-application/internal/pkg/errorswithcode"

var (
	ErrUpdateVersionStatus         = errorswithcode.NewValidationErrorWithCode("invalid status to update version", "INVALID_STATUS_TO_UPDATE")
	ErrInvalidStatusToSendToReview = errorswithcode.NewValidationErrorWithCode("invalid status to send to review", "INVALID_STATUS_TO_SEND_TO_REVIEW")
	ErrInvalidStatusToApprove      = errorswithcode.NewValidationErrorWithCode("invalid status to approve", "INVALID_STATUS_TO_APPROVE")
	ErrInvalidStatusToDecline      = errorswithcode.NewValidationErrorWithCode("invalid status to decline", "INVALID_STATUS_TO_DECLINE")
)

// Status is a type for a version status.
type Status string

const (
	DraftStatus    Status = "DRAFT"
	OnReviewStatus Status = "ON_REVIEW"
	ApprovedStatus Status = "APPROVED"
	DeclinedStatus Status = "DECLINED"
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
	case OnReviewStatus.String():
		return OnReviewStatus
	case ApprovedStatus.String():
		return ApprovedStatus
	case DeclinedStatus.String():
		return DeclinedStatus
	}
	panic("unknown status")
}

// IsDraft checks if a version status is draft.
func (s Status) IsDraft() bool {
	return s == DraftStatus
}

// IsOnReview checks if a version status is on review.
func (s Status) IsOnReview() bool {
	return s == OnReviewStatus
}

// IsApproved checks if a version status is approved.
func (s Status) IsApproved() bool {
	return s == ApprovedStatus
}

// IsDeclined checks if a version status is declined.
func (s Status) IsDeclined() bool {
	return s == DeclinedStatus
}

// sendToReview checks if a version status allows to send version to review.
// It returns OnReviewStatus.
func (s Status) sendToReview() (Status, error) {
	if s.IsDraft() {
		return OnReviewStatus, nil
	}
	return "", ErrInvalidStatusToSendToReview
}

// approve checks if a version status allows to approve version.
// It returns ApprovedStatus.
func (s Status) approve() (Status, error) {
	if s.IsOnReview() {
		return ApprovedStatus, nil
	}
	return "", ErrInvalidStatusToApprove
}

// decline checks if a version status allows to decline version and returns DeclinedStatus.
// It returns DeclinedStatus.
func (s Status) decline() (Status, error) {
	if s.IsOnReview() {
		return DeclinedStatus, nil
	}
	return "", ErrInvalidStatusToDecline
}

// canUpdate checks if a version status allows to update version.
func (s Status) canUpdate() error {
	if s.IsDraft() {
		return nil
	}
	return ErrUpdateVersionStatus
}
