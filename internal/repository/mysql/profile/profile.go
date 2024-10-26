package profile

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

func (r *Repository) GetUserByProfileCode(ctx context.Context, profileCode int) (profile entity.Profile, err error) {
	query := r.db.Table(profile.TableName()).Where("profile_code = ?", profileCode).Find(&profile)
	if err = query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}
	return
}
