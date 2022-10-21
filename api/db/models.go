package db

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name         string
	Type         string
	PurchaseDate string // RFC3339
	Description  string
	RecordID     uint
}

// A Record has one food
type Record struct {
	gorm.Model
	Food              Food
	Description       string
	EatingDate        string // RFC3339
	EatenQuantity     uint
	SatisfactionScore uint
	PhotoURL          string
}
