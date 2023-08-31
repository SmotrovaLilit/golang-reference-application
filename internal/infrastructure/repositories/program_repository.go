package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"reference-application/internal/application/interfaces/repositories"
	"reference-application/internal/domain/program"
)

var _ repositories.ProgramRepository = (*ProgramRepository)(nil)

// ProgramModel is a model for a program.
type ProgramModel struct {
	ID           string `gorm:"primaryKey"` // TODO change to uuid, fix in https://github.com/SmotrovaLilit/golang-reference-application/issues/11
	PlatformCode string
}

// TableName returns a table name.
func (ProgramModel) TableName() string { return "programs" }

// ProgramRepository is a repository to manage programs.
type ProgramRepository struct {
	db *gorm.DB
}

// NewProgramRepository creates a new ProgramRepository.
func NewProgramRepository(db *gorm.DB) *ProgramRepository {
	return &ProgramRepository{db: db}
}

// Save saves a program.
// This method is a part of ProgramRepository interface.
// It panics if an errors occurs.
func (r *ProgramRepository) Save(ctx context.Context, program program.Program) {
	err := r.db.WithContext(ctx).Save(&ProgramModel{
		ID:           program.ID().String(),
		PlatformCode: program.PlatformCode().String(),
	}).Error
	if err != nil {
		panic(err)
	}
}

// FindByID finds a program by id.
// This method is a part of ProgramRepository interface.
// It panics if an errors occurs.
func (r *ProgramRepository) FindByID(ctx context.Context, id program.ID) *program.Program {
	var model ProgramModel
	err := r.db.WithContext(ctx).First(&model, "id = ?", id.String()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}
	_p := program.NewProgram(
		program.MustNewID(model.ID),
		program.MustNewPlatformCode(model.PlatformCode),
	)
	return &_p
}
