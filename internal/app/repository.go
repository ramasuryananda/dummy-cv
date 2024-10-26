package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/tx"

	"gorm.io/gorm"
)

type Repositories struct {
	Profile profile.RepositoryProvider
	tx      tx.RepositoryProvider
	db      *gorm.DB
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		Profile: profile.New(db),
		tx:      tx.New(db),
		db:      db,
	}
}
