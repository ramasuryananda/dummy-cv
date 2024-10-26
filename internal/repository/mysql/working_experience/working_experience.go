package working_experience

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

func (r *Repository) SaveWorkingExperience(ctx context.Context, data entity.WorkingExperience) (err error) {
	result := r.db.Model(&data).Save(&data)
	if err = result.Error; err != nil {
		return
	}

	return
}

func (r *Repository) GetWorkingExperienceByProfileCode(ctx context.Context, code int) (workingExperience entity.WorkingExperience, err error) {
	result := r.db.Table(workingExperience.TableName()).Where("profile_code = ?", code).First(&workingExperience)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}

	return
}
