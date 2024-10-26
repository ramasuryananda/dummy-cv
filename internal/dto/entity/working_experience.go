package entity

import (
	"database/sql"
	"time"
)

type WorkingExperience struct {
	ID                uint64       `gorm:"primaryKey;autoIncrement"`
	ProfileCode       uint64       `gorm:"index:working_experience_index_2;not null"`
	WorkingExperience string       `gorm:"type:text;column:working_experience"`
	CreatedAt         time.Time    `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt         sql.NullTime `gorm:"autoUpdateTime;column:updated_at"`
}

func (WorkingExperience) TableName() string {
	return "working_experience"
}
