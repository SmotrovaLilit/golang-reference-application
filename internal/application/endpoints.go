package application

import (
	"github.com/google/wire"
	"reference-application/internal/application/commands/approveprogramversion"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/application/commands/declineprogramversion"
	"reference-application/internal/application/commands/sendtoreviewprogramversion"
	"reference-application/internal/application/commands/updateprogramversion"
)

type Endpoints struct {
	CreateProgramEndpoint              createprogram.Endpoint
	UpdateProgramVersionEndpoint       updateprogramversion.Endpoint
	SendToReviewProgramVersionEndpoint sendtoreviewprogramversion.Endpoint
	ApproveProgramVersionEndpoint      approveprogramversion.Endpoint
	DeclineProgramVersionEndpoint      declineprogramversion.Endpoint
}

var Set = wire.NewSet(
	wire.Struct(new(Endpoints), "*"),
)
