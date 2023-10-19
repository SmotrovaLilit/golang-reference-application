package updateprogramversion

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"log/slog"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/log"
	"reference-application/internal/pkg/optional"
	"reference-application/internal/pkg/resource"
)

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Command is a command to update a version.
type Command struct {
	id          version.ID
	name        version.Name
	description optional.Optional[version.Description]
	number      optional.Optional[version.Number]
}

// Handler is a handler to update a version.
type Handler struct {
	Repository repositories.VersionRepository
	Logger     *slog.Logger
}

// NewCommand is a constructor for Command.
func NewCommand(
	id version.ID,
	name version.Name,
	description optional.Optional[version.Description],
	number optional.Optional[version.Number],
) Command {
	return Command{
		id:          id,
		name:        name,
		description: description,
		number:      number,
	}
}

// Handle handles a command to update a version.
func (h Handler) Handle(ctx context.Context, cmd Command) error {
	ctx = resource.PopulateContextWithResourceID(ctx, cmd.id.String())
	logger := log.WithContext(ctx, h.Logger)
	_version := h.Repository.FindByID(ctx, cmd.id)
	if _version == nil {
		logger.Warn("program version not found")
		return sharederrors.ErrVersionNotFound
	}
	err := _version.Update(cmd.name, cmd.description, cmd.number)
	if err != nil {
		logger.Warn(fmt.Sprintf(
			"failed to update program version: %s", err.Error()))
		return err
	}
	h.Repository.Save(ctx, *_version)
	logger.Info("program version updated")
	return nil
}

// This need to make command endpoint to implement resource endpoint interface.
var _ resource.Endpoint = Endpoint(nil)

// Endpoint is an endpoint to update a version.
type Endpoint endpoint.Endpoint

// ResourceName returns the resource name.
// It uses for logging.
func (e Endpoint) ResourceName() string { return "programVersion" }

// ResourceAction returns the resource action.
// It uses for logging.
func (e Endpoint) ResourceAction() string { return "update" }

// NewEndpoint creates a new endpoint to update a version.
func NewEndpoint(handler Handler) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		cmd := request.(Command)
		err = handler.Handle(ctx, cmd)
		return nil, err
	}
}
