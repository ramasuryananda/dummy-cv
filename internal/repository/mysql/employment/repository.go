package employment

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetEmploymentByProfileCode(ctx context.Context, profileCode uint64) ([]entity.Employment, error)
	GetFirstEmploymentByProfileCodeandID(ctx context.Context, profileCode uint64, id uint64) (entity.Employment, error)
	CreateEmploymentData(ctx context.Context, data entity.Employment) (id uint64, err error)
	DeleteEmploymentData(ctx context.Context, profileCode uint64, id uint64) (err error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
