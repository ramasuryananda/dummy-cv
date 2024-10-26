package app

import (
	"github.com/ramasuryananda/dummy-cv/internal/handler/profile"
)

// Handlers types of handler layer.
type Handlers struct {
	Profile *profile.Handler
}

// New initializes handler layer.
func NewHandler(useCase *UseCases) *Handlers {
	return &Handlers{
		Profile: profile.New(useCase.Profile),
	}
}
