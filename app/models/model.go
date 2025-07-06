package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (mod *User) BeforeCreate(tx *gorm.DB) (err error) {
	mod.ID = uuid.NewString()
	return nil
}

func (mod *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	mod.ID = uuid.NewString()
	return nil
}
