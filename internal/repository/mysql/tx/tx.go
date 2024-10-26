package tx

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

// BeginTx starts a transaction.
// An error is returned if the begin transaction is not successful.
func (repository *Repository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*gorm.DB, error) {
	tx := repository.db.Begin(opts)
	if tx.Error != nil {
		return tx, tx.Error
	}

	return tx, nil
}

// Commit querying and running rows scan.
// An error is returned if the commit is not successful
func (repository *Repository) Commit(ctx context.Context, tx *gorm.DB) error {
	return tx.Commit().Error
}

// Rollback aborts the transaction.
// An error is returned if the rollback is not successful
func (repository *Repository) Rollback(ctx context.Context, tx *gorm.DB) error {
	return tx.Rollback().Error
}
