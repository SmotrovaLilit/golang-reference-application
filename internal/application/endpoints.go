package application

import (
	"github.com/google/wire"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/application/commands/updateprogramversion"
)

type Endpoints struct {
	CreateProgramEndpoint        createprogram.Endpoint
	UpdateProgramVersionEndpoint updateprogramversion.Endpoint
}

var Set = wire.NewSet(
	wire.Struct(new(Endpoints), "*"),
)
