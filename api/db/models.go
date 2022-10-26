package db

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name            string
	Type            string
	PurchaseDate    string // RFC3339
	CurrentQuantity float64
	Description     string
}

// A Record belongs to a food
type Record struct {
	gorm.Model
	FoodName          string
	Description       string
	EatingDate        string // RFC3339
	EatenQuantity     float64
	SatisfactionScore uint
	PhotoURL          string
}
