package request

type GetEducationRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
}

type CreateEducationRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	School      string `json:"school" validate:"required,max=100"`
	Degree      string `json:"degree" validate:"required,max=10"`
	StartDate   string `json:"startDate" validate:"required,date"`
	EndDate     string `json:"endDate" validate:"omitempty,date"`
	City        string `json:"city" validate:"required,max=50"`
	Description string `json:"description"`
}

type DeleteEducationRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	ID          uint64 `query:"id" validate:"required,numeric"`
}
