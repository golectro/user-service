package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Username  string         `gorm:"type:varchar(50);not null" json:"username"`
	Roles     datatypes.JSON `gorm:"type:json" json:"roles"`
	CreatedAt time.Time      `gorm:"type:datetime;column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"type:datetime;column:updated_at;autoUpdateTime:milli"`
}

func (User) TableName() string {
	return "users"
}
