package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string       `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt TimeWrapper  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt TimeWrapper  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt *TimeWrapper `json:"deletedAt,omitempty" gorm:"index"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}
