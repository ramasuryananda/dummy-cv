package education

import "github.com/ramasuryananda/dummy-cv/internal/usecase/education"

type Handler struct {
	educationUsecase education.UseCaseProvider
}

func New(educationUsecase education.UseCaseProvider) *Handler {
	return &Handler{
		educationUsecase: educationUsecase,
	}
}
