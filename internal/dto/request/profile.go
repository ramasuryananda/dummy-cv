package request

type GetUserProfileRequest struct {
	ProfileCode uint64 `param:"code" validate:"required,numeric"`
}

type CreateProfileRequest struct {
	WantedJobTitle string `json:"wantedJobTitle"  validate:"required,max=255"`
	FirstName      string `json:"firstName" validate:"required,max=50"`
	LastName       string `json:"lastName" validate:"omitempty,max=50"`
	Email          string `json:"email"  validate:"omitempty,max=50"`
	Phone          string `json:"phone"  validate:"omitempty,max=15"`
	Country        string `json:"country"  validate:"omitempty,max=20"`
	City           string `json:"city"  validate:"omitempty,max=20"`
	Address        string `json:"address" `
	PostalCode     uint64 `json:"postalCode"  validate:"omitempty,numeric"`
	DrivingLicense string `json:"drivingLicense"  validate:"omitempty,max=30"`
	Nationality    string `json:"nationality"  validate:"required,max=20"`
	PlaceOfBirth   string `json:"placeOfBirth"  validate:"required,max=20"`
	DateOfBirth    string `json:"dateOfBirth"  validate:"required,date"`
}

type UpdateProfileRequest struct {
	ProfileCode    uint64 `param:"code" validate:"required,numeric"`
	WantedJobTitle string `json:"wantedJobTitle"  validate:"required,max=255"`
	FirstName      string `json:"firstName"  validate:"required,max=50"`
	LastName       string `json:"lastName"  validate:"omitempty,max=50"`
	Email          string `json:"email"  validate:"omitempty,max=50"`
	Phone          string `json:"phone"  validate:"omitempty,max=15"`
	Country        string `json:"country"  validate:"omitempty,max=20"`
	City           string `json:"city"  validate:"omitempty,max=20"`
	Address        string `json:"address" `
	PostalCode     uint64 `json:"postalCode"  validate:"omitempty,numeric"`
	DrivingLicense string `json:"drivingLicense"  validate:"omitempty,max=30"`
	Nationality    string `json:"nationality"  validate:"required,max=20"`
	PlaceOfBirth   string `json:"placeOfBirth"  validate:"required,max=20"`
	DateOfBirth    string `json:"dateOfBirth" validate:"required,date"`
}
