package employment

import "github.com/ramasuryananda/dummy-cv/internal/usecase/employment"

type Handler struct {
	employmentUsecase employment.UseCaseProvider
}

func New(employmentUsecase employment.UseCaseProvider) *Handler {
	return &Handler{
		employmentUsecase: employmentUsecase,
	}
}
