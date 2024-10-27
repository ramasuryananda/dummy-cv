package skill

import (
	"context"

	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/skill"
)

type UseCaseProvider interface {
	GetUserSkill(ctx context.Context, req request.GetSkillRequest) (resp writer.Response, status int)
	CreateSkillData(ctx context.Context, req request.CreateSkillRequest) (resp writer.Response, status int)
	DeleteSkillData(ctx context.Context, req request.DeleteSkillRequest) (resp writer.Response, status int)
}

type UseCase struct {
	skillRepository   skill.RepositoryProvider
	profileRepository profile.RepositoryProvider
}

func New(skillRepository skill.RepositoryProvider, profileRepository profile.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		skillRepository:   skillRepository,
		profileRepository: profileRepository,
	}
}
