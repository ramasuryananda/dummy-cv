package entity

import (
	"database/sql"
	"time"
)

type Education struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement;not null;column:id"`
	ProfileCode uint64       `gorm:"not null;index:employment_index_3;column:profile_code"`
	School      string       `gorm:"type:varchar(100);not null;column:school"`
	Degree      string       `gorm:"type:varchar(10);not null;column:degree"`
	StartDate   time.Time    `gorm:"type:date;not null;column:start_date"`
	EndDate     sql.NullTime `gorm:"type:date;column:end_date"`
	City        string       `gorm:"type:varchar(50);not null;column:city"`
	Description string       `gorm:"type:text;column:description"`
	CreatedAt   time.Time    `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   sql.NullTime `gorm:"autoUpdateTime;column:updated_at"`
}

func (Education) TableName() string {
	return "education"
}
