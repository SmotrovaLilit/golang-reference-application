package readmodels

import (
	"github.com/google/wire"
	"reference-application/internal/application/queries/approvedprograms"
)

var Set = wire.NewSet(
	NewApprovedProgramsReadModel,
	wire.Bind(new(approvedprograms.ReadModel), new(*ApprovedProgramsReadModel)),
)
