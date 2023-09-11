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
		optional.Of(version.MustNewDescription("description")),
		optional.Of(version.MustNewNumber("1.0.0")),
	)
	return _version, _program
}

// NewPreparedToReviewVersion returns a new draft version prepared to review.
func NewPreparedToReviewVersion() (version.Version, program.Program) {
	_program := program.NewProgram(
		program.MustNewID("d942802b-02a2-4ac2-ab78-28f89aed2994"),
		program.IPhonePlatformCode,
	)
	_version := version.NewExistingVersion(
		version.MustNewID("b2082806-3d32-48d4-9fe4-ce67b22be21e"),
		"ios-smart-calculator",
		_program.ID(),
		version.DraftStatus,
		optional.Of(version.MustNewDescription("description")),
		optional.Of(version.MustNewNumber("1.0.0")),
	)
	return _version, _program
}

// NewProgram returns a new program.
func NewProgram(id program.ID) program.Program {
	return program.NewProgram(
		id,
		program.AndroidPlatformCode,
	)
}

// NewDraftVersionInProgram returns a new draft version in the program.
func NewDraftVersionInProgram(program program.Program, versionID version.ID) version.Version {
	name := version.MustNewName("ios-smart-calculator" + versionID.String())
	return version.NewExistingVersion(
		versionID,
		name,
		program.ID(),
		version.DraftStatus,
		optional.Empty[version.Description](),
		optional.Empty[version.Number](),
	)
}

// NewApprovedVersionInProgram returns a new approved version in the program.
func NewApprovedVersionInProgram(program program.Program, versionID version.ID) version.Version {
	name := version.MustNewName("ios-smart-calculator" + versionID.String())
	return version.NewExistingVersion(
		versionID,
		name,
		program.ID(),
		version.ApprovedStatus,
		optional.Of(version.MustNewDescription("description")),
		optional.Of(version.MustNewNumber("1.0.0")),
	)
}

// NewOnReviewVersionInProgram returns a new version on review in the program.
func NewOnReviewVersionInProgram(program program.Program, versionID version.ID) version.Version {
	name := version.MustNewName("ios-smart-calculator" + versionID.String())
	return version.NewExistingVersion(
		versionID,
		name,
		program.ID(),
		version.OnReviewStatus,
		optional.Of(version.MustNewDescription("description")),
		optional.Of(version.MustNewNumber("1.0.0")),
	)
}
