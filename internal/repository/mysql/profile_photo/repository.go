package profile_photo

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetUserProfilePhotoByProfileCode(ctx context.Context, profileCode int) (profilePhoto entity.ProfilePhoto, err error)
	SaveUserProfilePhoto(ctx context.Context, profilePhoto entity.ProfilePhoto) (err error)
	DeleteUserProfilePhoto(ctx context.Context, profilePhoto entity.ProfilePhoto) (err error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
