package request

type UpsertPhotoProfileRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
	Base64Image string `json:"base64img" validate:"required"`
}

type DownloadPhotoProfileRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
}

type DeletePhotoProfileRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
}
