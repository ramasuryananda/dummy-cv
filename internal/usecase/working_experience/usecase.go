package working_experience

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/working_experience"
)

type UseCaseProvider interface {
	UpsertUserWorkingExperience(ctx context.Context, req request.UpsertWorkingExperienceRequest) (resp writer.Response, httpStatus int)
	GetUserWorkingExperience(ctx context.Context, req request.GetUserWorkingExperienceRequest) (resp writer.Response, httpStatus int)
}

type UseCase struct {
	workingExperienceRepository working_experience.RepositoryProvider
	profileRepository           profile.RepositoryProvider
}

func New(workingExperienceRepository working_experience.RepositoryProvider, profileRepository profile.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		workingExperienceRepository: workingExperienceRepository,
		profileRepository:           profileRepository,
	}
}
