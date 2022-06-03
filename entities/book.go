package entities

import "time"

type Book struct {
	ID          uint64     `json:"id" gorm:"primarykey" example:"1"`
	Title       *string    `json:"title" gorm:"type:varchar(255)" example:"abc"`
	Description *string    `json:"description" gorm:"type:varchar(255)" example:"abc"`
	Content     *string    `json:"content" gorm:"type:text" example:"abc"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;autoCreateTime" example:"1970-01-01T00:00:00Z"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;autoUpdateTime" example:"1970-01-01T00:00:00Z"`
}

func (Book) TableName() string {
	return "book"
}
