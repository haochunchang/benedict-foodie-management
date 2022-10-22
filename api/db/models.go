package db

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name            string
	Type            string
	PurchaseDate    string // RFC3339
	CurrentQuantity uint
	Description     string
}

// A Record belongs to a food
type Record struct {
	gorm.Model
	FoodID            uint
	Food              Food
	Description       string
	EatingDate        string // RFC3339
	EatenQuantity     uint
	SatisfactionScore uint
	PhotoURL          string
}
