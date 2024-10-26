package profile

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"gorm.io/gorm"
)

func (r *Repository) GetUserByProfileCode(ctx context.Context, profileCode int) (profile entity.Profile, err error) {
	err = r.db.Table(profile.TableName()).Where("profile_code = ?", profileCode).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}
	return
}

func (r *Repository) InsertProfile(ctx context.Context, profileData entity.Profile) (profileCode uint64, err error) {
	result := r.db.Model(&profileData).Create(&profileData)
	if err = result.Error; err != nil {
		return
	}

	profileCode = profileData.ProfileCode

	return
}

func (r *Repository) UpdateProfile(ctx context.Context, profileData entity.Profile) (profileCode uint64, err error) {
	result := r.db.Model(&profileData).Updates(&profileData)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorDatabaseNotFound
		}
		return
	}

	profileCode = profileData.ProfileCode

	return
}
