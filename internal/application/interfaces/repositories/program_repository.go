package repositories

import (
	"context"
	"reference-application/internal/domain/program"
)

// ProgramRepository is a repository to save a program.
type ProgramRepository interface {
	// Save saves a program.
	Save(ctx context.Context, program program.Program)
}
