package employment

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/employment"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
)

type UseCaseProvider interface {
	GetUserEmployment(ctx context.Context, req request.GetEmploymentRequest) (resp writer.Response, status int)
	CreateEmploymentData(ctx context.Context, req request.CreateEmploymentRequest) (resp writer.Response, status int)
	DeleteEmploymentData(ctx context.Context, req request.DeleteEmploymentRequest) (resp writer.Response, status int)
}

type UseCase struct {
	employmentRepository employment.RepositoryProvider
	profileRepository    profile.RepositoryProvider
}

func New(employmentRepository employment.RepositoryProvider, profileRepository profile.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		employmentRepository: employmentRepository,
		profileRepository:    profileRepository,
	}
}
