package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	//ID uint `gorm:"primaryKey" json:"id"`
	//CreatedAt time.Time      `json:"created_at"`
	//UpdatedAt time.Time      `json:"updated_at"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	ProductID uint   `json:"product_id"` // Relationship
	Comment   string `gorm:"not null" json:"comment"`
	Stars     uint   `gorm:"not null" json:"stars"`
}
