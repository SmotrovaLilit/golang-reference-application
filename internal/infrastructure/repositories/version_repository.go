package repositories

import (
	"context"
	"gorm.io/gorm"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
)

var _ repositories.VersionRepository = (*VersionRepository)(nil)

// VersionModel is a model for a version.
type VersionModel struct {
	ID        string       `gorm:"primaryKey"` // TODO change to uuid, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/11
	Name      string       // TODO set limit for length
	ProgramID string       // TODO change to uuid, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/11
	Program   ProgramModel `gorm:"foreignKey:ProgramID;references:ID"`
	Status    string       // TODO change to enum
}

// TableName returns a table name.
func (VersionModel) TableName() string { return "versions" }

// VersionRepository is a repository to manage versions.
type VersionRepository struct {
	db *gorm.DB
}

// NewVersionRepository creates a new VersionRepository.
func NewVersionRepository(db *gorm.DB) *VersionRepository {
	return &VersionRepository{db}
}

// Save saves a version.
// This method is a part of VersionRepository interface.
// It panics if an errors occurs.
func (r *VersionRepository) Save(ctx context.Context, version version.Version) {
	err := r.db.WithContext(ctx).Save(&VersionModel{
		ID:        version.ID().String(),
		Name:      version.Name().String(),
		ProgramID: version.ProgramID().String(),
		Status:    version.Status().String(),
	}).Error
	if err != nil {
		panic(err)
	}
}

// FindByID finds a version by id.
// This method is a part of VersionRepository interface.
// It panics if an errors occurs.
func (r *VersionRepository) FindByID(ctx context.Context, id version.ID) version.Version {
	var model VersionModel
	err := r.db.WithContext(ctx).First(&model, "id = ?", id.String()).Error
	if err != nil {
		panic(err)
	}
	return version.NewExistingVersion(
		version.MustNewID(model.ID),
		version.MustNewName(model.Name),
		program.MustNewID(model.ProgramID),
		version.MustNewStatus(model.Status),
	)
}
