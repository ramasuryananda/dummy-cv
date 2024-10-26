package app

import (
	profile_photo "github.com/ramasuryananda/dummy-cv/internal/handler/photo"
	"github.com/ramasuryananda/dummy-cv/internal/handler/profile"
	"github.com/ramasuryananda/dummy-cv/internal/handler/working_experience"
)

// Handlers types of handler layer.
type Handlers struct {
	Profile           *profile.Handler
	ProfilePhoto      *profile_photo.Handler
	WorkingExperience *working_experience.Handler
}

// New initializes handler layer.
func NewHandler(useCase *UseCases) *Handlers {
	return &Handlers{
		Profile:           profile.New(useCase.Profile),
		ProfilePhoto:      profile_photo.New(useCase.PhotoProfile),
		WorkingExperience: working_experience.New(useCase.WorkingExperience),
	}
}
