package updateprogramversion

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/errors"
)

// ErrVersionNotFound is a version not found errors.
var ErrVersionNotFound = errors.New("version not found", "NOT_FOUND")

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Command is a command to update a version.
type Command struct {
	ID   version.ID
	name version.Name
}

// Handler is a handler to update a version.
type Handler struct {
	Repository repositories.VersionRepository
}

// NewCommand is a constructor for Command.
func NewCommand(id version.ID, name version.Name) Command {
	return Command{
		ID:   id,
		name: name,
	}
}

// Handle handles a command to update a version.
func (h Handler) Handle(ctx context.Context, cmd Command) error {
	_version := h.Repository.FindByID(ctx, cmd.ID)
	if _version == nil {
		return ErrVersionNotFound
	}
	err := _version.UpdateName(cmd.name)
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
