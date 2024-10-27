package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/usecase/education"
	"github.com/ramasuryananda/dummy-cv/internal/usecase/employment"
	profile_photo "github.com/ramasuryananda/dummy-cv/internal/usecase/photo"
	"github.com/ramasuryananda/dummy-cv/internal/usecase/profile"
	"github.com/ramasuryananda/dummy-cv/internal/usecase/working_experience"
)

type UseCases struct {
	Profile           profile.UseCaseProvider
	PhotoProfile      profile_photo.UseCaseProvider
	WorkingExperience working_experience.UseCaseProvider
	Employment        employment.UseCaseProvider
	Education         education.UseCaseProvider
}

// NewUseCase initializes useCase layer.
func NewUseCase(repositories *Repositories) *UseCases {
	return &UseCases{
		Profile:           profile.New(repositories.Profile),
		PhotoProfile:      profile_photo.New(repositories.Profile, repositories.ProfilePhoto),
		WorkingExperience: working_experience.New(repositories.WorkingExperience, repositories.Profile),
		Employment:        employment.New(repositories.EmploymentRepository, repositories.Profile),
		Education:         education.New(repositories.EducationRepository, repositories.Profile),
	}
}
