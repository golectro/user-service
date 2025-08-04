package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Email    string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"type:datetime;column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"type:datetime;column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (User) TableName() string {
	return "users"
}
