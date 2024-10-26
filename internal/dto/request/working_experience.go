package request

type UpsertWorkingExperienceRequest struct {
	ProfileCode       int    `param:"code" validate:"required,numeric"`
	WorkingExperience string `json:"workingExperience" validate:"required,min=1"`
}

type GetUserWorkingExperienceRequest struct {
	ProfileCode int `param:"code" validate:"required,numeric"`
}
