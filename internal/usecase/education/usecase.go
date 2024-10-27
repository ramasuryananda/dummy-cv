package education

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/education"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
)

type UseCaseProvider interface {
	GetUserEducation(ctx context.Context, req request.GetEducationRequest) (resp writer.Response, status int)
	CreateEducationData(ctx context.Context, req request.CreateEducationRequest) (resp writer.Response, status int)
	DeleteEducationData(ctx context.Context, req request.DeleteEducationRequest) (resp writer.Response, status int)
}

type UseCase struct {
	educationRepository education.RepositoryProvider
	profileRepository   profile.RepositoryProvider
}

func New(educationRepository education.RepositoryProvider, profileRepository profile.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		educationRepository: educationRepository,
		profileRepository:   profileRepository,
	}
}
