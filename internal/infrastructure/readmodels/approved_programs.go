package readmodels

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"reference-application/internal/application/queries/approvedprograms"
	"reference-application/internal/domain/version"
	"reference-application/internal/pkg/pager"
	"reference-application/internal/pkg/slices"
)

// ApprovedProgramsReadModel is a read model for programs query.
// It implements programs.ReadModel interface.
type ApprovedProgramsReadModel struct {
	db *gorm.DB
}

// NewApprovedProgramsReadModel is a constructor for ApprovedProgramsReadModel.
func NewApprovedProgramsReadModel(db *gorm.DB) *ApprovedProgramsReadModel {
	return &ApprovedProgramsReadModel{db: db}
}

// Query queries programs with total count.
// It implements programs.ReadModel interface.
func (m ApprovedProgramsReadModel) Query(ctx context.Context, pager pager.Pager) approvedprograms.Result {
	type versionDBModel struct {
		ID          string `gorm:"column:version_id"`
		Name        string `gorm:"column:name"`
		Number      string `gorm:"column:number"`
		Description string `gorm:"column:description"`
		Status      string `gorm:"column:status"`
	}

	type programDBModel struct {
		ID           string         `gorm:"column:id"`
		PlatformCode string         `gorm:"column:platform_code"`
		Version      versionDBModel `gorm:"embedded"`
	}

	var dbPrograms []programDBModel

	err := m.db.WithContext(ctx).
		Table("programs").
		Offset(pager.Offset().Int()).
		Limit(pager.Limit().Int()).
		Joins(
			"inner join (?) v ON programs.id = v.program_id",
			m.db.Raw(fmt.Sprintf(`SELECT *,
           		ROW_NUMBER() OVER (PARTITION BY program_id) AS number_in_concrete_program
    			FROM versions
    			WHERE status='%s'`, version.ApprovedStatus)),
		).
		Where("v.number_in_concrete_program = 1 OR v.number_in_concrete_program IS NULL").
		Select("programs.id as id, platform_code, v.id as version_id, name, number, description, status").
		Scan(&dbPrograms).Error
	if err != nil {
		panic(err)
	}

	return slices.Convert(func(dbProgram programDBModel) approvedprograms.Program {
		return approvedprograms.Program{
			ID:           dbProgram.ID,
			PlatformCode: dbProgram.PlatformCode,
			Version:      approvedprograms.Version(dbProgram.Version),
		}
	}, dbPrograms)
}
