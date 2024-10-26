package profile

import (
	"github.com/ramasuryananda/dummy-cv/internal/repository/mysql/profile"
)

type UseCaseProvider interface {
}

type UseCase struct {
	profileRepository profile.RepositoryProvider
}

func New(profileRepository profile.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		profileRepository: profileRepository,
	}
}
