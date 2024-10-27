package request

type GetEmploymentRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
}

type CreateEmploymentRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	JobTitle    string `json:"jobTitle" validate:"required,max=50"`
	Employer    string `json:"employer" validate:"required,max=50"`
	StartDate   string `json:"startDate" validate:"required,date"`
	EndDate     string `json:"endDate" validate:"omitempty,date"`
	City        string `json:"city" validate:"required,max=50"`
	Description string `json:"description"`
}

type DeleteEmploymentRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	ID          uint64 `query:"id" validate:"required,numeric"`
}
