package profile

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
)

type UseCaseProvider interface {
	GetUserProfile(ctx context.Context, code uint64) (resp writer.Response, httpStatus int)
	CreateUserProfile(ctx context.Context, req request.CreateProfileRequest) (resp writer.Response, httpStatus int)
	UpdateUserProfile(ctx context.Context, req request.UpdateProfileRequest) (resp writer.Response, httpStatus int)
}

type UseCase struct {
	profileRepository profile.RepositoryProvider
}

func New(profileRepository profile.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		profileRepository: profileRepository,
	}
}
