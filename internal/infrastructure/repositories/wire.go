package repositories

import (
	"github.com/google/wire"
	"reference-application/internal/application/interfaces/repositories"
)

var Set = wire.NewSet(
	NewProgramRepository,
	NewVersionRepository,
	wire.Bind(new(repositories.ProgramRepository), new(*ProgramRepository)),
	wire.Bind(new(repositories.VersionRepository), new(*VersionRepository)),
)
