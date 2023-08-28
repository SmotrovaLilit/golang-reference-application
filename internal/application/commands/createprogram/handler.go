package createprogram

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
)

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Command is a command to create a program.
type Command struct {
	ID           program.ID
	PlatformCode program.PlatformCode
}

// NewCommand is a constructor for Command.
func NewCommand(id program.ID, platformCode program.PlatformCode) Command {
	return Command{
		ID:           id,
		PlatformCode: platformCode,
	}
}

// Handler is a handler to create a program.
type Handler struct {
	Repository repositories.ProgramRepository
}

// Handle handles a command to create a program.
func (h Handler) Handle(ctx context.Context, cmd Command) {
	_program := program.NewProgram(cmd.ID, cmd.PlatformCode)
	h.Repository.Save(ctx, _program)
}

// Endpoint is an endpoint to create a program.
type Endpoint endpoint.Endpoint

// NewEndpoint creates a new endpoint to create a program.
func NewEndpoint(handler Handler) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		command := request.(Command)
		handler.Handle(ctx, command)
		return nil, nil
	}
}
