package response

import (
	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
)

type EmploymentDataResponse struct {
	ID          uint64          `json:"id"`
	JobTitle    string          `json:"jobTitle"`
	Employer    string          `json:"employer"`
	StartDate   general.YMDDate `json:"startDate"`
	EndDate     general.YMDDate `json:"endDate"`
	City        string          `json:"city"`
	Description string          `json:"description"`
}

type CreateEmploymentResponse struct {
	ProfileCode uint64 `json:"profileCode"`
	Id          uint64 `json:"id"`
}

type DeleteEmploymentResponse struct {
	ProfileCode uint64 `json:"profileCode"`
}
