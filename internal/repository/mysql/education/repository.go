package education

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetEducationByProfileCode(ctx context.Context, profileCode uint64) ([]entity.Education, error)
	GetFirstEducationByProfileCodeandID(ctx context.Context, profileCode uint64, id uint64) (entity.Education, error)
	CreateEducationData(ctx context.Context, data entity.Education) (id uint64, err error)
	DeleteEducationData(ctx context.Context, profileCode uint64, id uint64) (err error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
