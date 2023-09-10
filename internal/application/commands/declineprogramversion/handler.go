package declineprogramversion

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/application/sharederrors"
	"reference-application/internal/domain/version"
)

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Command is a command to update a version.
type Command struct {
	id version.ID
}

// Handler is a handler to update a version.
type Handler struct {
	Repository repositories.VersionRepository
}

// NewCommand is a constructor for Command.
func NewCommand(
	id version.ID,
) Command {
	return Command{
		id: id,
	}
}

// Handle handles a command to update a version.
func (h Handler) Handle(ctx context.Context, cmd Command) error {
	_version := h.Repository.FindByID(ctx, cmd.id)
	if _version == nil {
		return sharederrors.ErrVersionNotFound
	}
	err := _version.Decline()
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
