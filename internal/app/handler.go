package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/handler/education"
	"github.com/ramasuryananda/dummy-cv/internal/handler/employment"
	profile_photo "github.com/ramasuryananda/dummy-cv/internal/handler/photo"
	"github.com/ramasuryananda/dummy-cv/internal/handler/profile"
	"github.com/ramasuryananda/dummy-cv/internal/handler/skill"
	"github.com/ramasuryananda/dummy-cv/internal/handler/working_experience"
)

// Handlers types of handler layer.
type Handlers struct {
	Profile           *profile.Handler
	ProfilePhoto      *profile_photo.Handler
	WorkingExperience *working_experience.Handler
	Employment        *employment.Handler
	Education         *education.Handler
	Skill             *skill.Handler
}

// New initializes handler layer.
func NewHandler(useCase *UseCases) *Handlers {
	return &Handlers{
		Profile:           profile.New(useCase.Profile),
		ProfilePhoto:      profile_photo.New(useCase.PhotoProfile),
		WorkingExperience: working_experience.New(useCase.WorkingExperience),
		Employment:        employment.New(useCase.Employment),
		Education:         education.New(useCase.Education),
		Skill:             skill.New(useCase.Skill),
	}
}
