package app

import (
	profile_photo "github.com/ramasuryananda/dummy-cv/internal/handler/photo"
	"github.com/ramasuryananda/dummy-cv/internal/handler/profile"
)

// Handlers types of handler layer.
type Handlers struct {
	Profile      *profile.Handler
	ProfilePhoto *profile_photo.Handler
}

// New initializes handler layer.
func NewHandler(useCase *UseCases) *Handlers {
	return &Handlers{
		Profile:      profile.New(useCase.Profile),
		ProfilePhoto: profile_photo.New(useCase.PhotoProfile),
	}
}
