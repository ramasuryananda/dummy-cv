package profile_photo

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile_photo"
)

type UseCaseProvider interface {
	UpsertUserPhotoProfile(ctx context.Context, req request.UpsertPhotoProfileRequest) (resp writer.Response, httpStatus int)
	DownloadPhotoProfile(ctx context.Context, req request.DownloadPhotoProfileRequest) (resp writer.Response, httpStatus int)
	DeletePhotoProfile(ctx context.Context, req request.DeletePhotoProfileRequest) (resp writer.Response, httpStatus int)
}

type UseCase struct {
	profileRepository      profile.RepositoryProvider
	profilePhotoRepository profile_photo.RepositoryProvider
}

func New(profileRepository profile.RepositoryProvider, profilePhotoRepository profile_photo.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		profileRepository:      profileRepository,
		profilePhotoRepository: profilePhotoRepository,
	}
}
