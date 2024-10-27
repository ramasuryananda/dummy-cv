package education

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

func (r *Repository) GetEducationByProfileCode(ctx context.Context, profileCode uint64) (data []entity.Education, err error) {
	err = r.db.Table(entity.Education{}.TableName()).Where("profile_code = ?", profileCode).Find(&data).Error
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetFirstEducationByProfileCodeandID(ctx context.Context, profileCode uint64, id uint64) (data entity.Education, err error) {
	err = r.db.Table(entity.Education{}.TableName()).Where(map[string]interface{}{
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

func (r *Repository) CreateEducationData(ctx context.Context, data entity.Education) (id uint64, err error) {
	err = r.db.Model(&data).Create(&data).Error
	if err != nil {
		return
	}

	id = data.ID
	return
}

func (r *Repository) DeleteEducationData(ctx context.Context, profileCode uint64, id uint64) (err error) {
	err = r.db.Table(entity.Education{}.TableName()).Where(map[string]interface{}{
		"profile_code": profileCode,
		"id":           id,
	}).Delete(entity.Education{}).Error
	if err != nil {
		return
	}

	return
}
