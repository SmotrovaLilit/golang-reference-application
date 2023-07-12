package program

import "reference-application/internal/pkg/errors"

// ErrInvalidPlatformCode is an errors for invalid platform code.
var ErrInvalidPlatformCode = errors.New("invalid program platform code", "INVALID_PROGRAM_PLATFORM_CODE")

const (
	// AndroidPlatformCode is a constant for android platform code.
	AndroidPlatformCode PlatformCode = "ANDROID"

	// IPhonePlatformCode is a constant for iphone platform code.
	IPhonePlatformCode PlatformCode = "IPHONE"
)

// PlatformCode is a type for platform code.
type PlatformCode string

// NewPlatformCode is a constructor for PlatformCode.
func NewPlatformCode(raw string) (PlatformCode, error) {
	code := PlatformCode(raw)
	switch code {
	case AndroidPlatformCode, IPhonePlatformCode:
		return code, nil
	default:
		return "", ErrInvalidPlatformCode
	}
}

// String returns a string representation of PlatformCode.
func (code PlatformCode) String() string { return string(code) }
