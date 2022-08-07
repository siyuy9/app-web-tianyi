package entity

import (
	"time"

	"gorm.io/gorm"
)

// copy of gorm.Model with added json annotations
type commonFields struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
