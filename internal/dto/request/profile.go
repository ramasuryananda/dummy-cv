package request

type GetUserProfileRequest struct {
	ProfileCode int `param:"code" validate:"required,numeric"`
}

type CreateProfileRequest struct {
	WantedJobTitle string `json:"wantedJobTitle" name:"Wanted Job Title" validate:"required,max=255"`
	FirstName      string `json:"firstName" name:"First Name" validate:"required,max=50"`
	LastName       string `json:"lastName" name:"Last Name" validate:"omitempty,max=50"`
	Email          string `json:"email" name:"Email" validate:"omitempty,max=50"`
	Phone          string `json:"phone" name:"Phone" validate:"omitempty,max=15"`
	Country        string `json:"country" name:"Country" validate:"omitempty,max=20"`
	City           string `json:"city" name:"City" validate:"omitempty,max=20"`
	Address        string `json:"address" name:"Address"`
	PostalCode     uint64 `json:"postalCode" name:"Postal Code" validate:"omitempty,numeric"`
	DrivingLicense string `json:"drivingLicense" name:"Driving License Number" validate:"omitempty,max=30"`
	Nationality    string `json:"nationality" name:"Nationality" validate:"required,max=20"`
	PlaceOfBirth   string `json:"placeOfBirth" name:"Place Of Birth" validate:"required,max=20"`
	DateOfBirth    string `json:"dateOfBirth" name:"Date Of Birth" validate:"required,date"`
}

type UpdateProfileRequest struct {
	ProfileCode    int    `param:"code" validate:"required,numeric"`
	WantedJobTitle string `json:"wantedJobTitle" name:"Wanted Job Title" validate:"required,max=255"`
	FirstName      string `json:"firstName" name:"First Name" validate:"required,max=50"`
	LastName       string `json:"lastName" name:"Last Name" validate:"omitempty,max=50"`
	Email          string `json:"email" name:"Email" validate:"omitempty,max=50"`
	Phone          string `json:"phone" name:"Phone" validate:"omitempty,max=15"`
	Country        string `json:"country" name:"Country" validate:"omitempty,max=20"`
	City           string `json:"city" name:"City" validate:"omitempty,max=20"`
	Address        string `json:"address" name:"Address"`
	PostalCode     uint64 `json:"postalCode" name:"Postal Code" validate:"omitempty,numeric"`
	DrivingLicense string `json:"drivingLicense" name:"Driving License Number" validate:"omitempty,max=30"`
	Nationality    string `json:"nationality" name:"Nationality" validate:"required,max=20"`
	PlaceOfBirth   string `json:"placeOfBirth" name:"Place Of Birth" validate:"required,max=20"`
	DateOfBirth    string `json:"dateOfBirth" name:"Date Of Birth" validate:"required,date"`
}
