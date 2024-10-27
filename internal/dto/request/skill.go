package request

type GetSkillRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
}

type CreateSkillRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	Skill       string `json:"skill" validate:"required,max=20"`
	Level       string `json:"level" validate:"required,max=20"`
}

type DeleteSkillRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	ID          uint64 `query:"id" validate:"required,numeric"`
}
