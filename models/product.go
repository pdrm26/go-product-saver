package models

import (
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Product   string    `json:"product"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
