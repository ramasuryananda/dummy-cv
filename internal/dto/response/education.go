package response

import (
	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
)

type EducationDataResponse struct {
	ID          uint64          `json:"id"`
	School      string          `json:"school"`
	Degree      string          `json:"degree"`
	StartDate   general.YMDDate `json:"startDate"`
	EndDate     general.YMDDate `json:"endDate"`
	City        string          `json:"city"`
	Description string          `json:"description"`
}

type CreateEducationResponse struct {
	ProfileCode uint64 `json:"profileCode"`
	Id          uint64 `json:"id"`
}

type DeleteEducationResponse struct {
	ProfileCode uint64 `json:"profileCode"`
}
