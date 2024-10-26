package working_experience

import (
	"github.com/ramasuryananda/dummy-cv/internal/usecase/working_experience"
)

type Handler struct {
	workingExperienceUsecase working_experience.UseCaseProvider
}

func New(workingExperienceUsecase working_experience.UseCaseProvider) *Handler {
	return &Handler{
		workingExperienceUsecase: workingExperienceUsecase,
	}
}
