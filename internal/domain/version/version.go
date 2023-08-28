package version

import "reference-application/internal/domain/program"

// Version is a type for a program version.
type Version struct {
	id        ID
	name      Name
	programID program.ID
	status    Status
}

// NewVersion is a constructor for Version.
func NewVersion(id ID, name Name, programID program.ID) Version {
	return Version{
		id:        id,
		name:      name,
		programID: programID,
		status:    DraftStatus,
	}
}

// NewExistingVersion is a constructor for a already existing Version.
func NewExistingVersion(
	id ID,
	name Name,
	programID program.ID,
	status Status) Version {
	return Version{
		id:        id,
		name:      name,
		programID: programID,
		status:    status,
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
