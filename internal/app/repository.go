package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile_photo"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/working_experience"

	"gorm.io/gorm"
)

type Repositories struct {
	Profile           profile.RepositoryProvider
	ProfilePhoto      profile_photo.RepositoryProvider
	WorkingExperience working_experience.RepositoryProvider
	db                *gorm.DB
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		Profile:           profile.New(db),
		ProfilePhoto:      profile_photo.New(db),
		WorkingExperience: working_experience.New(db),
		db:                db,
	}
}
