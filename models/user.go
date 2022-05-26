package models

import "time"

type User struct {
	ID           uint64     `json:"id" gorm:"primarykey" example:"1"`
	Username     *string    `json:"username" gorm:"uniqueIndex;type:varchar(32)" example:"abc"`
	Email        *string    `json:"email" gorm:"uniqueIndex;type:varchar(255)" example:"abc@gmail.com"`
	PasswordHash *string    `json:"passwordHash" gorm:"column:password_hash;type:varchar(64)" example:""`
	Role         *string    `json:"role" gorm:"type:varchar(8)" example:"ADMIN"`
	Status       *string    `json:"status" gorm:"type:varchar(16)" example:"ACTIVE"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;autoCreateTime" example:"1970-01-01T00:00:00Z"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;autoUpdateTime" example:"1970-01-01T00:00:00Z"`
}

func (User) TableName() string {
	return "user"
}
