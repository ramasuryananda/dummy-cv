package profile_photo

import (
	profile_photo "github.com/ramasuryananda/dummy-cv/internal/usecase/photo"
)

type Handler struct {
	photoProfileUsecase profile_photo.UseCaseProvider
}

func New(profilePhotoUsecase profile_photo.UseCaseProvider) *Handler {
	return &Handler{
		photoProfileUsecase: profilePhotoUsecase,
	}
}
