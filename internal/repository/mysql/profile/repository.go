package profile

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetUserByProfileCode(ctx context.Context, profileCode uint64) (entity.Profile, error)
	InsertProfile(ctx context.Context, profileData entity.Profile) (profileCode uint64, err error)
	UpdateProfile(ctx context.Context, profileData entity.Profile) (profileCode uint64, err error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
