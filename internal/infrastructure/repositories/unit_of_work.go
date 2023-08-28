package repositories

import (
	"context"
	"gorm.io/gorm"
	"reference-application/internal/application/interfaces/repositories"
)

// UnitOfWork is a transaction wrapper.
type UnitOfWork struct {
	db *gorm.DB
}

// NewUnitOfWork creates a new UnitOfWork.
func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

// Do executes a function in a transaction.
func (u UnitOfWork) Do(ctx context.Context, f func(repositories.UnitOfWorkStore)) {
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	f(unitOfWorkStore{tx: tx})
	err := tx.Commit().Error
	if err != nil {
		panic(err)
	}
}

// unitOfWorkStore is a store for transaction repositories.
type unitOfWorkStore struct {
	tx                *gorm.DB
	programRepository *ProgramRepository
	versionRepository *VersionRepository
}

// ProgramRepository returns a ProgramRepository in a transaction context.
func (u unitOfWorkStore) ProgramRepository() repositories.ProgramRepository {
	if u.programRepository != nil {
		return u.programRepository
	}
	u.programRepository = NewProgramRepository(u.tx)
	return u.programRepository
}

// VersionRepository returns a VersionRepository in a transaction context.
func (u unitOfWorkStore) VersionRepository() repositories.VersionRepository {
	if u.versionRepository != nil {
		return u.versionRepository
	}
	u.versionRepository = NewVersionRepository(u.tx)
	return u.versionRepository
}
