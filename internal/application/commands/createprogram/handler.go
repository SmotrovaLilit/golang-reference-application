package createprogram

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
)

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Command is a command to create a program.
type Command struct {
	ID           program.ID // TODO make it private
	PlatformCode program.PlatformCode
	VersionID    version.ID
	VersionName  version.Name
}

// NewCommand is a constructor for Command.
func NewCommand(
	id program.ID,
	platformCode program.PlatformCode,
	versionID version.ID,
	versionName version.Name,
) Command {
	return Command{
		ID:           id,
		PlatformCode: platformCode,
		VersionID:    versionID,
		VersionName:  versionName,
	}
}

// Handler is a handler to create a program.
type Handler struct {
	Repository        repositories.ProgramRepository
	VersionRepository repositories.VersionRepository
}

// Handle handles a command to create a program.
func (h Handler) Handle(ctx context.Context, cmd Command) {
	_program := program.NewProgram(cmd.ID, cmd.PlatformCode)
	_version := version.NewVersion(cmd.VersionID, cmd.VersionName, _program.ID())
	// TODO start transaction
	h.Repository.Save(ctx, _program)
	h.VersionRepository.Save(ctx, _version)
	// End transaction
}

// Endpoint is an endpoint to create a program.
type Endpoint endpoint.Endpoint

// NewEndpoint creates a new endpoint to create a program.
// TODO no one test with this function, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/10
func NewEndpoint(handler Handler) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		command := request.(Command)
		handler.Handle(ctx, command)
		return nil, nil
	}
}
