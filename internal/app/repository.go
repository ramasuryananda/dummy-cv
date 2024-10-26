package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile_photo"

	"gorm.io/gorm"
)

type Repositories struct {
	Profile      profile.RepositoryProvider
	ProfilePhoto profile_photo.RepositoryProvider
	db           *gorm.DB
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		Profile:      profile.New(db),
		ProfilePhoto: profile_photo.New(db),
		db:           db,
	}
}
