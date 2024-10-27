package working_experience

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	SaveWorkingExperience(ctx context.Context, data entity.WorkingExperience) (err error)
	GetWorkingExperienceByProfileCode(ctx context.Context, code uint64) (workingExperience entity.WorkingExperience, err error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
