package skill

import (
	"github.com/ramasuryananda/dummy-cv/internal/usecase/skill"
)

type Handler struct {
	skillUsecase skill.UseCaseProvider
}

func New(skillUsecase skill.UseCaseProvider) *Handler {
	return &Handler{
		skillUsecase: skillUsecase,
	}
}
