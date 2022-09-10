package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	//ID uint `gorm:"primaryKey" json:"id"`
	//CreatedAt time.Time      `json:"created_at"`
	//UpdatedAt time.Time      `json:"updated_at"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Title       string   `gorm:"not null" json:"title"`
	Description string   `gorm:"not null" json:"description"`
	Price       float32  `gorm:"not null" json:"price"`
	Reviews     []Review `gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductID" json:"reviews,omitempty"`
}
