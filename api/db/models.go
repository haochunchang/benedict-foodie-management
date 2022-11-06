package db

import (
	"time"

	"gorm.io/gorm"
)

// A food has many eating Record.
type Food struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Type        string
	Description string
	Records     []Record `gorm:"foreignKey:FoodName;references:Name;constraint:OnUpdate:CASCADE;"`
}

type Record struct {
	gorm.Model
	FoodName          string
	Description       string
	EatingDate        time.Time
	EatenQuantity     float64
	SatisfactionScore uint
	PhotoURL          string
}
