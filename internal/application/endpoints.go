package application

import (
	"github.com/google/wire"
	"reference-application/internal/application/commands/createprogram"
)

type Endpoints struct {
	CreateProgramEndpoint createprogram.Endpoint
}

var Set = wire.NewSet(
	wire.Struct(new(Endpoints), "*"),
)
