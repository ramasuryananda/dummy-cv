package profile

import (
	"github.com/ramasuryananda/dummy-cv/internal/usecase/profile"
)

type Handler struct {
	profileUsecase profile.UseCaseProvider
}

func New(profileUsecase profile.UseCaseProvider) *Handler {
	return &Handler{
		profileUsecase: profileUsecase,
	}
}
