package db

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name            string
	Type            string
	PurchaseDate    time.Time
	NutritionFactID uint
}

type Record struct {
	gorm.Model
	FoodID            uint
	Food              Food
	Description       string
	EatingDate        time.Time
	EatenQuantity     uint
	SatisfactionScore uint
	PhotoURL          string
}
