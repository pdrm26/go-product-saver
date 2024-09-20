package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         int     `json:"id"`
	Product    string  `json:"product"`
	Price      float64 `json:"price"`
	Category   Category `gorm:"foriegnKey:CategoryID"`
	CategoryID string
	Brand      Brand `gorm:"foriegnKey:BrandID"`
	BrandID    string
	CreatedAt  time.Time `json:"created_at"`
}

type Brand struct {
	gorm.Model
	ID    int    `json:"id"`
	Brand string `json:"brand"`
}

type Category struct {
	gorm.Model
	ID       int    `json:"id"`
	Category string `json:"category"`
}
