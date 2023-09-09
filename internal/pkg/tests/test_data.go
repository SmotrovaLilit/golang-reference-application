package tests

import (
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/optional"
)

// NewDraftVersion returns a new draft version.
func NewDraftVersion() (version.Version, program.Program) {
	_program := program.NewProgram(
		program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
		program.AndroidPlatformCode,
	)
	_version := version.NewVersion(
		version.MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
		"smart-calculator",
		_program.ID(),
	)
	return _version, _program
}

// NewOnReviewVersion returns a new version on review.
func NewOnReviewVersion() (version.Version, program.Program) {
	_program := program.NewProgram(
		program.MustNewID("d942802b-02a2-4ac2-ab78-28f89aed2994"),
		program.IPhonePlatformCode,
	)
	_version := version.NewExistingVersion(
		version.MustNewID("b2082806-3d32-48d4-9fe4-ce67b22be21e"),
		"ios-smart-calculator",
		_program.ID(),
		version.OnReviewStatus,
		optional.Empty[version.Description](),
		optional.Empty[version.Number](),
	)
	return _version, _program
}
