package db

import (
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connectionInfo string) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Food{})
	db.AutoMigrate(&Record{})
}

func GetDB(connectionInfo string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CreateFood(db *gorm.DB, food Food) error {
	if _, err := time.Parse(time.RFC3339, food.PurchaseDate); err != nil {
		return err
	}
	return db.Create(&food).Error
}

func GetFoodByName(db *gorm.DB, name string) Food {
	var result Food
	db.Where("name = ?", name).Find(&result)
	return result
}

func CreateRecord(db *gorm.DB, r Record) error {
	if _, err := time.Parse(time.RFC3339, r.EatingDate); err != nil {
		return err
	}
	if _, err := time.Parse(time.RFC3339, r.Food.PurchaseDate); err != nil {
		return err
	}
	return db.Create(&r).Error
}

func GetRecordsByDate(db *gorm.DB, eatingDate string) ([]Record, error) {
	var results []Record
	_, err := time.Parse(time.RFC3339, eatingDate)
	if err != nil {
		return results, err
	}
	start := strings.Split(eatingDate, "T")[0] + "T00:00:00+08:00"
	end := strings.Split(eatingDate, "T")[0] + "T23:59:59+08:00"

	query := db.Model(&Record{}).Preload("Food")
	query.Where("eating_date BETWEEN ? AND ?", start, end).Find(&results)
	return results, nil
}
