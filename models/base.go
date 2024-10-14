package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Check out omitzero proposal: https://github.com/golang/go/issues/45669

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *BaseModel) BeforeCreate(tx *gorm.DB) error {
	var err error
	if t.ID == uuid.Nil {
		t.ID, err = uuid.NewV7()
	}

	return err
}
