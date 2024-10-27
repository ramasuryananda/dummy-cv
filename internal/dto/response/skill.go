package response

type SkillDataResponse struct {
	ID    uint64 `json:"id"`
	Skill string `json:"skill"`
	Level string `json:"level"`
}

type CreateSkillResponse struct {
	ProfileCode uint64 `json:"profileCode"`
	Id          uint64 `json:"id"`
}

type DeleteSkillResponse struct {
	ProfileCode uint64 `json:"profileCode"`
}
