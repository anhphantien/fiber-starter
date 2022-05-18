package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint64         `json:"id" gorm:"column:id;primarykey"`
	Title     *string        `json:"title" gorm:"type:varchar(255)"`
	Author    *string        `json:"author" gorm:"type:varchar(255)"`
	Publisher *string        `json:"publisher" gorm:"type:varchar(255)"`
	Name      *string        `json:"name" gorm:"type:varchar(255)"`
	CreatedAt *time.Time     `json:"createdAt" gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt *time.Time     `json:"updatedAt" gorm:"column:updated_at;type:datetime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:datetime;autoDeleteTime"`
}

func (Book) TableName() string {
	return "book"
}
