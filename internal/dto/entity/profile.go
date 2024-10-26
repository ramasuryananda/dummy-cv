package entity

import (
	"database/sql"
	"time"
)

type Profile struct {
	ProfileCode    uint64       `gorm:"primaryKey;autoIncrement;column:profile_code"`
	WantedJobTitle string       `gorm:"size:255;not null;column:wanted_job_title"`
	FirstName      string       `gorm:"size:50;not null;column:first_name"`
	LastName       string       `gorm:"size:50;column:last_name"`
	Email          string       `gorm:"size:50;column:email"`
	Phone          string       `gorm:"size:15;column:phone"`
	Country        string       `gorm:"size:20;column:country"`
	City           string       `gorm:"size:20;column:city"`
	Address        string       `gorm:"type:text;column:address"`
	PostalCode     uint64       `gorm:"column:postal_code"`
	DrivingLicense string       `gorm:"size:30;column:driving_license"`
	Nationality    string       `gorm:"size:20;not null;column:nationality"`
	PlaceOfBirth   string       `gorm:"size:30;not null;column:place_of_birth"`
	DateOfBirth    time.Time    `gorm:"type:date;not null;column:date_of_birth"`
	CreatedAt      time.Time    `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      sql.NullTime `gorm:"autoUpdateTime;column:updated_at"`
}

// TableName overrides the table name used by Gorm.
func (Profile) TableName() string {
	return "profile"
}
