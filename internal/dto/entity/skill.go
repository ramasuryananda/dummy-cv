package entity

import (
	"database/sql"
	"time"
)

type Skill struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement;not null;column:id"`
	ProfileCode uint64       `gorm:"not null;index:employment_index_3;column:profile_code"`
	Skill       string       `gorm:"type:varchar(20);not null;column:skill"`
	Level       string       `gorm:"type:varchar(20);not null;column:level"`
	CreatedAt   time.Time    `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   sql.NullTime `gorm:"autoUpdateTime;column:updated_at"`
}

func (Skill) TableName() string {
	return "skill"
}
