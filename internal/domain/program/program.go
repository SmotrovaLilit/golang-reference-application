package program

// Program is a type for program.
type Program struct {
	id           ID
	platformCode PlatformCode
}

// NewProgram creates a new Program.
func NewProgram(id ID, platformCode PlatformCode) Program {
	return Program{id: id, platformCode: platformCode}
}

// PlatformCode returns a platform code.
func (p Program) PlatformCode() PlatformCode { return p.platformCode }

// ID returns an ID.
func (p Program) ID() ID { return p.id }
