package app

import (
	profile_photo "github.com/ramasuryananda/dummy-cv/internal/usecase/photo"
	"github.com/ramasuryananda/dummy-cv/internal/usecase/profile"
)

type UseCases struct {
	Profile      profile.UseCaseProvider
	PhotoProfile profile_photo.UseCaseProvider
}

// NewUseCase initializes useCase layer.
func NewUseCase(repositories *Repositories) *UseCases {
	return &UseCases{
		Profile:      profile.New(repositories.Profile),
		PhotoProfile: profile_photo.New(repositories.Profile, repositories.ProfilePhoto),
	}
}
