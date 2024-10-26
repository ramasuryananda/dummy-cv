package request

type UpsertPhotoProfileRequest struct {
	ProfileCode int    `param:"code" validate:"required,numeric"`
	Base64Image string `json:"base64img" validate:"required"`
}

type DownloadPhotoProfileRequest struct {
	ProfileCode int `param:"code" validate:"required,numeric"`
}

type DeletePhotoProfileRequest struct {
	ProfileCode int `param:"code" validate:"required,numeric"`
}
