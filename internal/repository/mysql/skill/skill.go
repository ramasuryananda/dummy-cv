package skill

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

func (r *Repository) GetSkillByProfileCode(ctx context.Context, profileCode uint64) (data []entity.Skill, err error) {
	err = r.db.Table(entity.Skill{}.TableName()).Where("profile_code = ?", profileCode).Find(&data).Error
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetFirstSkillByProfileCodeandID(ctx context.Context, profileCode uint64, id uint64) (data entity.Skill, err error) {
	err = r.db.Table(entity.Skill{}.TableName()).Where(map[string]interface{}{
		"profile_code": profileCode,
		"id":           id,
	}).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}
	return
}

func (r *Repository) CreateSkillData(ctx context.Context, data entity.Skill) (id uint64, err error) {
	err = r.db.Model(&data).Create(&data).Error
	if err != nil {
		return
	}

	id = data.ID
	return
}

func (r *Repository) DeleteSkillData(ctx context.Context, profileCode uint64, id uint64) (err error) {
	err = r.db.Table(entity.Skill{}.TableName()).Where(map[string]interface{}{
		"profile_code": profileCode,
		"id":           id,
	}).Delete(entity.Skill{}).Error
	if err != nil {
		return
	}

	return
}
