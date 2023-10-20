package createprogram

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"log/slog"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/log"
	"reference-application/internal/pkg/resource"
)

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Command is a command to create a program.
type Command struct {
	id           program.ID
	platformCode program.PlatformCode
	versionID    version.ID
	versionName  version.Name
}

// NewCommand is a constructor for Command.
func NewCommand(
	id program.ID,
	platformCode program.PlatformCode,
	versionID version.ID,
	versionName version.Name,
) Command {
	return Command{
		id:           id,
		platformCode: platformCode,
		versionID:    versionID,
		versionName:  versionName,
	}
}

// Handler is a handler to create a program.
type Handler struct {
	UnitOfWork repositories.UnitOfWork
	Logger     *slog.Logger
}

// Handle handles a command to create a program.
func (h Handler) Handle(ctx context.Context, cmd Command) {
	ctx = resource.PopulateContextWithResourceID(ctx, cmd.id.String())
	logger := log.WithContext(ctx, h.Logger)
	_program := program.NewProgram(cmd.id, cmd.platformCode)
	_version := version.NewVersion(cmd.versionID, cmd.versionName, _program.ID())

	h.UnitOfWork.Do(ctx, func(store repositories.UnitOfWorkStore) {
		store.ProgramRepository().Save(ctx, _program)
		store.VersionRepository().Save(ctx, _version)
	})
	logger.Info("program created")
}

var _ resource.Endpoint = Endpoint(nil)

// Endpoint is an endpoint to update a version.
type Endpoint endpoint.Endpoint

// ResourceName returns the resource name.
// It uses for logging.
func (e Endpoint) ResourceName() string { return "program" }

// ResourceAction returns the resource action.
// It uses for logging.
func (e Endpoint) ResourceAction() string { return "create" }

// NewEndpoint creates a new endpoint to create a program.
func NewEndpoint(handler Handler) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		command := request.(Command)
		handler.Handle(ctx, command)
		return nil, nil
	}
}
