package response

type UpdatePhotoProfileResponse struct {
	ProfileCode uint64 `json:"profileCode"`
	PhotoURL    string `json:"photo_url"`
}

type DowndloadPhotoProfile struct {
	Base64String string
	PhotoURL     string
}

type DeletePhotoProfileResponse struct {
	ProfileCode uint64 `json:"profileCode"`
}
