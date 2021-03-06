package entities

import "time"

type User struct {
	ID             uint64     `json:"id"             gorm:"column:id;              primarykey"                                      example:"1"`
	Username       *string    `json:"username"       gorm:"column:username;        type:varchar(32);  uniqueIndex"                  example:"abc"`
	Email          *string    `json:"email"          gorm:"column:email;           type:varchar(255); uniqueIndex"                  example:"abc@gmail.com"`
	HashedPassword *string    `json:"-"              gorm:"column:hashed_password; type:varchar(64)"                                example:""`
	Role           *string    `json:"role"           gorm:"column:role;            type:varchar(8)"                                 example:"ADMIN"`
	Status         *string    `json:"status"         gorm:"column:status;          type:varchar(16)"                                example:"ACTIVE"`
	CreatedAt      *time.Time `json:"createdAt"      gorm:"column:created_at;      type:datetime;     autoCreateTime"               example:"1970-01-01T00:00:00Z"`
	UpdatedAt      *time.Time `json:"updatedAt"      gorm:"column:updated_at;      type:datetime;     autoUpdateTime"               example:"1970-01-01T00:00:00Z"`
	Books          []Book     `json:"books"          gorm:"foreignKey:UserID;      references:ID;     constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "user"
}
