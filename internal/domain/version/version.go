package version

import (
	"reference-application/internal/domain/program"
	. "reference-application/internal/pkg/optional"
)

// Version is a type for a program version.
type Version struct {
	id          ID
	name        Name
	programID   program.ID
	status      Status
	description Optional[Description]
}

// NewVersion is a constructor for Version.
func NewVersion(id ID, name Name, programID program.ID) Version {
	return Version{
		id:          id,
		name:        name,
		programID:   programID,
		status:      DraftStatus,
		description: Empty[Description](),
	}
}

// NewExistingVersion is a constructor for a already existing Version.
func NewExistingVersion(
	id ID,
	name Name,
	programID program.ID,
	status Status,
	description Optional[Description],
) Version {
	return Version{
		id:          id,
		name:        name,
		programID:   programID,
		status:      status,
		description: description,
	}
}

// ID returns a version ID.
func (v *Version) ID() ID {
	return v.id
}

// Name returns a version name.
func (v *Version) Name() Name {
	return v.name
}

// ProgramID returns a version program ID.
func (v *Version) ProgramID() program.ID {
	return v.programID
}

// Status returns a version status.
func (v *Version) Status() Status {
	return v.status
}

// Description returns a version description.
func (v *Version) Description() Optional[Description] {
	return v.description
}

// Update updates a version.
func (v *Version) Update(name Name, description Optional[Description]) error {
	if err := v.canUpdate(); err != nil {
		return err
	}
	v.name = name
	v.description = description
	return nil
}

// canUpdate checks if a version can be updated.
func (v *Version) canUpdate() error {
	return v.status.canUpdate()
}
