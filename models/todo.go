package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	BaseModel
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	IsCompleted bool      `json:"is_completed" gorm:"type:boolean;not null"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}
