package application

import (
	"context"
	"github.com/google/wire"
	"log/slog"
	"reference-application/internal/application/commands/approveprogramversion"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/application/commands/declineprogramversion"
	"reference-application/internal/application/commands/sendtoreviewprogramversion"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/application/queries/approvedprograms"
	"reference-application/internal/pkg/panicrecovery"
)

var Set = wire.NewSet(
	NewEndpoints,
)

type Endpoints struct {
	CreateProgramEndpoint              createprogram.Endpoint
	UpdateProgramVersionEndpoint       updateprogramversion.Endpoint
	SendToReviewProgramVersionEndpoint sendtoreviewprogramversion.Endpoint
	ApproveProgramVersionEndpoint      approveprogramversion.Endpoint
	DeclineProgramVersionEndpoint      declineprogramversion.Endpoint
	ApprovedProgramsEndpoint           approvedprograms.Endpoint
}

type endpointType interface {
	~func(ctx context.Context, request interface{}) (response interface{}, err error)
}

// getEndpointMiddlewares returns endpoint middlewares.
func getEndpointMiddlewares[T endpointType](e T, logger *slog.Logger) T {
	return panicrecovery.EndpointMiddleware(e, logger)
}

// NewEndpoints is a constructor for Endpoints.
func NewEndpoints(
	logger *slog.Logger,
	createProgramEndpoint createprogram.Endpoint,
	updateProgramVersionEndpoint updateprogramversion.Endpoint,
	sendToReviewProgramVersionEndpoint sendtoreviewprogramversion.Endpoint,
	approveProgramVersionEndpoint approveprogramversion.Endpoint,
	declineProgramVersionEndpoint declineprogramversion.Endpoint,
	approvedProgramsEndpoint approvedprograms.Endpoint,
) Endpoints {
	return Endpoints{
		CreateProgramEndpoint:              getEndpointMiddlewares(createProgramEndpoint, logger),
		UpdateProgramVersionEndpoint:       getEndpointMiddlewares(updateProgramVersionEndpoint, logger),
		SendToReviewProgramVersionEndpoint: getEndpointMiddlewares(sendToReviewProgramVersionEndpoint, logger),
		ApproveProgramVersionEndpoint:      getEndpointMiddlewares(approveProgramVersionEndpoint, logger),
		DeclineProgramVersionEndpoint:      getEndpointMiddlewares(declineProgramVersionEndpoint, logger),
		ApprovedProgramsEndpoint:           getEndpointMiddlewares(approvedProgramsEndpoint, logger),
	}
}
