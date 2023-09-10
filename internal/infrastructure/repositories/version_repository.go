package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/optional"
)

var _ repositories.VersionRepository = (*VersionRepository)(nil)

// VersionModel is a model for a version.
type VersionModel struct {
	ID          string       `gorm:"primaryKey"` // TODO change to uuid, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/11
	Name        string       // TODO set limit for length
	ProgramID   string       // TODO change to uuid, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/11
	Program     ProgramModel `gorm:"foreignKey:ProgramID;references:id"`
	Status      string       // TODO change to enum
	Description string       // TODO set limit for length
	Number      string       // TODO set limit for length
}

// TableName returns a table name.
func (VersionModel) TableName() string { return "versions" }

// VersionRepository is a repository to manage versions.
type VersionRepository struct {
	db *gorm.DB
}

// NewVersionRepository creates a new VersionRepository.
func NewVersionRepository(db *gorm.DB) *VersionRepository {
	return &VersionRepository{db: db}
}

// Save saves a version.
// This method is a part of VersionRepository interface.
// It panics if an errors occurs.
func (r *VersionRepository) Save(ctx context.Context, version version.Version) {
	descriptionValue := ""
	if version.Description().IsPresent() {
		descriptionValue = version.Description().Value().String()
	}
	numberValue := ""
	if version.Number().IsPresent() {
		numberValue = version.Number().Value().String()
	}
	err := r.db.WithContext(ctx).Save(&VersionModel{
		ID:          version.ID().String(),
		Name:        version.Name().String(),
		ProgramID:   version.ProgramID().String(),
		Status:      version.Status().String(),
		Description: descriptionValue,
		Number:      numberValue,
	}).Error
	if err != nil {
		panic(err)
	}
}

// FindByID finds a version by id.
// This method is a part of VersionRepository interface.
// It returns nil if a version is not found.
// It panics if an errors occurs.
func (r *VersionRepository) FindByID(ctx context.Context, id version.ID) *version.Version {
	var model VersionModel
	err := r.db.WithContext(ctx).First(&model, "id = ?", id.String()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}
	description := optional.Empty[version.Description]()
	if model.Description != "" {
		description = optional.Of[version.Description](version.MustNewDescription(model.Description))
	}
	number := optional.Empty[version.Number]()
	if model.Number != "" {
		number = optional.Of[version.Number](version.MustNewNumber(model.Number))
	}

	_v := version.NewExistingVersion(
		version.MustNewID(model.ID),
		version.MustNewName(model.Name),
		program.MustNewID(model.ProgramID),
		version.MustNewStatus(model.Status),
		description,
		number,
	)
	return &_v
}
