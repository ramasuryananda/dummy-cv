package response

import "github.com/ramasuryananda/dummy-cv/internal/dto/general"

type GetProfileResponse struct {
	WantedJobTitle string          `json:"wantedJobTitle"`
	FirstName      string          `json:"firstName"`
	LastName       string          `json:"lastName"`
	Email          string          `json:"email"`
	Phone          string          `json:"phone"`
	Country        string          `json:"country"`
	City           string          `json:"city"`
	Address        string          `json:"address"`
	PostalCode     uint64          `json:"postalCode"`
	DrivingLicense string          `json:"drivingLicense"`
	Nationality    string          `json:"nationality"`
	PlaceOfBirth   string          `json:"placeOfBirth"`
	DateOfBirth    general.YMDDate `json:"dateOfBirth"`
}

type CreateProfileResponse struct {
	ProfileCode uint64 `json:"profileCode"`
}
