package healthcheck

import (
	"context"
	"database/sql"
	"fmt"
)

// NewDatabaseChecker returns a checker that checks connection to DB.
func NewDatabaseChecker(database *sql.DB) CheckerFunc {
	return func(ctx context.Context) error {
		if database == nil {
			return fmt.Errorf("database is nil")
		}

		return database.PingContext(ctx)
	}
}
