package skill

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetSkillByProfileCode(ctx context.Context, profileCode uint64) ([]entity.Skill, error)
	GetFirstSkillByProfileCodeandID(ctx context.Context, profileCode uint64, id uint64) (entity.Skill, error)
	CreateSkillData(ctx context.Context, data entity.Skill) (id uint64, err error)
	DeleteSkillData(ctx context.Context, profileCode uint64, id uint64) (err error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
