package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/usecase/profile"
)

type UseCases struct {
	Profile profile.UseCaseProvider
}

// NewUseCase initializes useCase layer.
func NewUseCase(repositories *Repositories) *UseCases {
	return &UseCases{
		Profile: profile.New(repositories.Profile),
	}
}
