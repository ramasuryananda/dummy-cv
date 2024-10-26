package entity

import (
	"database/sql"
	"time"
)

type ProfilePhoto struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement;column:id"`
	ProfileCode uint64       `gorm:"not null;column:profile_code"`
	PhotoURL    string       `gorm:"type:text;column:photo_url"`
	CreatedAt   time.Time    `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   sql.NullTime `gorm:"autoUpdateTime;column:updated_at"`
}

// TableName overrides the table name used by Gorm.
func (ProfilePhoto) TableName() string {
	return "profile_photo"
}
