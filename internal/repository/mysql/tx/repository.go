package tx

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

// RepositoryProvider represents client repository methods.
type RepositoryProvider interface {
	// BeginTx starts a transaction.
	// An error is returned if the begin transaction is not successful.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*gorm.DB, error)

	// Commit querying and running rows scan.
	// An error is returned if the commit is not successful
	Commit(ctx context.Context, tx *gorm.DB) error

	// Rollback aborts the transaction.
	// An error is returned if the rollback is not successful
	Rollback(ctx context.Context, tx *gorm.DB) error
}

// Repository types of repository layer.
type Repository struct {
	db *gorm.DB
}

// New initializes repository layer.
func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
