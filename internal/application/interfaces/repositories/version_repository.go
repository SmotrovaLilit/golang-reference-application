package repositories

import (
	"context"
	"reference-application/internal/domain/version"
)

// VersionRepository is a repository to save a program versions.
type VersionRepository interface {
	// Save saves a program.
	Save(ctx context.Context, version version.Version)
}
