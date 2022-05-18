package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint64         `json:"id" gorm:"primarykey"`
	Title     *string        `json:"title" gorm:"type:varchar(255)"`
	Author    *string        `json:"author" gorm:"type:varchar(255)"`
	Publisher *string        `json:"publisher" gorm:"type:varchar(255)"`
	Name      *string        `json:"name" gorm:"type:varchar(255)"`
	CreatedAt *time.Time     `json:"created_at" gorm:"type:datetime(0);autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"type:datetime(0);autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime(0);autoDeleteTime"`
}
