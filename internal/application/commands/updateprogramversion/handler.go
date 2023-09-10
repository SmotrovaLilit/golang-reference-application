package updateprogramversion

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/optional"
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
	_version := h.Repository.FindByID(ctx, cmd.id)
	if _version == nil {
		return sharederrors.ErrVersionNotFound
	}
	err := _version.Update(cmd.name, cmd.description, cmd.number)
	if err != nil {
		return err
	}
	h.Repository.Save(ctx, *_version)
	return nil
}

// Endpoint is an endpoint to update a version.
type Endpoint endpoint.Endpoint

// NewEndpoint creates a new endpoint to update a version.
func NewEndpoint(handler Handler) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		cmd := request.(Command)
		err = handler.Handle(ctx, cmd)
		return nil, err
	}
}
