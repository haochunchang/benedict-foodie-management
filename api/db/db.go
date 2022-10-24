package db

import (
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB returns a database handle specified by connectionInfo
func GetDB(connectionInfo string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type Repository interface {
	Init()
	Clear()
}

// FoodRepoistory stores the food information
type FoodRepository interface {
	Repository
	CreateFood(food Food) error
	GetFoodByName(name string) Food
}

type FoodRepositoryPSQL struct {
	db *gorm.DB
}

func NewFoodRepositoryPSQL(conn *gorm.DB) *FoodRepositoryPSQL {
	return &FoodRepositoryPSQL{conn}
}

func (f *FoodRepositoryPSQL) Init() {
	f.db.AutoMigrate(&Food{})
}

func (f *FoodRepositoryPSQL) Clear() {
	f.db.Exec("DROP TABLE IF EXISTS foods")
}

func (f *FoodRepositoryPSQL) CreateFood(food Food) error {
	if _, err := time.Parse(time.RFC3339, food.PurchaseDate); err != nil {
		return err
	}
	return f.db.Create(&food).Error
}

func (f *FoodRepositoryPSQL) GetFoodByName(name string) Food {
	var result Food
	f.db.Where("name = ?", name).Find(&result)
	return result
}

// RecordRepoistory stores the food record data
type RecordRepository interface {
	Repository
	CreateRecord(r Record) error
	GetRecordsByDate(eatingDate string) ([]Record, error)
}

type RecordRepositoryPSQL struct {
	db *gorm.DB
}

func NewRecordRepositoryPSQL(conn *gorm.DB) *RecordRepositoryPSQL {
	return &RecordRepositoryPSQL{conn}
}

func (rr *RecordRepositoryPSQL) Init() {
	rr.db.AutoMigrate(&Record{})
}

func (rr *RecordRepositoryPSQL) Clear() {
	rr.db.Exec("DROP TABLE IF EXISTS records")
}

func (rr *RecordRepositoryPSQL) CreateRecord(r Record) error {
	if _, err := time.Parse(time.RFC3339, r.EatingDate); err != nil {
		return err
	}
	return rr.db.Create(&r).Error
}

func (rr *RecordRepositoryPSQL) GetRecordsByDate(eatingDate string) ([]Record, error) {
	var results []Record
	_, err := time.Parse(time.RFC3339, eatingDate)
	if err != nil {
		return results, err
	}
	start := strings.Split(eatingDate, "T")[0] + "T00:00:00+08:00"
	end := strings.Split(eatingDate, "T")[0] + "T23:59:59+08:00"

	query := rr.db.Model(&Record{}).Preload("Food")
	query.Where("eating_date BETWEEN ? AND ?", start, end).Find(&results)
	return results, nil
}
