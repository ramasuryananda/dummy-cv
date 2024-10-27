package entity

import (
	"database/sql"
	"time"
)

type Employment struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement;not null;column:id"`
	ProfileCode uint64       `gorm:"not null;index:employment_index_3;column:profile_code"`
	JobTitle    string       `gorm:"type:varchar(50);not null;column:job_title"`
	Employer    string       `gorm:"type:varchar(50);not null;column:employer"`
	StartDate   time.Time    `gorm:"type:date;not null;column:start_date"`
	EndDate     sql.NullTime `gorm:"type:date;column:end_date"`
	City        string       `gorm:"type:varchar(50);not null;column:city"`
	Description string       `gorm:"type:text;column:description"`
	CreatedAt   time.Time    `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   sql.NullTime `gorm:"autoUpdateTime;column:updated_at"`
}

func (Employment) TableName() string {
	return "employment"
}
