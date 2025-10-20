package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string  `gorm:"size:150;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
}
