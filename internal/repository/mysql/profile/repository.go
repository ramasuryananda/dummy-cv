package profile

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetUserByProfileCode(ctx context.Context, profileCode int) (entity.Profile, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
