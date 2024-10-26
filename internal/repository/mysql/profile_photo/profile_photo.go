package profile_photo

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

func (r *Repository) GetUserProfilePhotoByProfileCode(ctx context.Context, profileCode int) (profilePhoto entity.ProfilePhoto, err error) {
	result := r.db.Table(profilePhoto.TableName()).Where("profile_code = ?", profileCode).First(&profilePhoto)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}

	return
}

func (r *Repository) SaveUserProfilePhoto(ctx context.Context, profilePhoto entity.ProfilePhoto) (err error) {
	result := r.db.Model(&profilePhoto).Save(&profilePhoto)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}

	return
}

func (r *Repository) DeleteUserProfilePhoto(ctx context.Context, profilePhoto entity.ProfilePhoto) (err error) {
	result := r.db.Model(&profilePhoto).Delete(&profilePhoto)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}

	return
}
