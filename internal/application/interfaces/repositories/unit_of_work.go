package repositories

import "context"

// UnitOfWorkStore is a store for transaction repositories.
type UnitOfWorkStore interface {
	ProgramRepository() ProgramRepository
	VersionRepository() VersionRepository
}

// UnitOfWork is a transaction wrapper.
type UnitOfWork interface {
	Do(context.Context, func(UnitOfWorkStore))
}
