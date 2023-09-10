package repositories

import (
	"context"
	"reference-application/internal/domain/version"
)

// VersionRepository is a repository to save a program versions.
type VersionRepository interface {
	// Save saves a version.
	Save(ctx context.Context, version version.Version)
	// FindByID finds a version by id.
	FindByID(ctx context.Context, id version.ID) *version.Version
}
