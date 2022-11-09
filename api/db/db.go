package db

import (
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
	CreateFood(Food) error
	GetFoodByName(string) (Food, error)
	UpdateFoodByName(string, Food) error
	DeleteFood(Food) error
	CreateRecord(Record) error
	GetRecordsByDate(int64, int64, int64) ([]Record, error)
	UpdateRecord(Record, Record) error
	DeleteRecord(Record) error
}

type FoodRepositoryPSQL struct {
	db *gorm.DB
}

func NewFoodRepositoryPSQL(conn *gorm.DB) *FoodRepositoryPSQL {
	return &FoodRepositoryPSQL{conn}
}

func (f *FoodRepositoryPSQL) Init() {
	f.db.AutoMigrate(&Food{}, &Record{})
}

func (f *FoodRepositoryPSQL) Clear() {
	f.db.Exec("DROP TABLE IF EXISTS records")
	f.db.Exec("DROP TABLE IF EXISTS foods")
}

func (f *FoodRepositoryPSQL) CreateFood(food Food) error {
	return f.db.Create(&food).Error
}

func (f *FoodRepositoryPSQL) GetFoodByName(name string) (result Food, err error) {
	err = f.db.Where("name = ?", name).Last(&result).Error
	return
}

func (f *FoodRepositoryPSQL) UpdateFoodByName(name string, food Food) error {
	var oldFood Food
	if err := f.db.Where("name = ?", name).Find(&oldFood).Error; err != nil {
		return err
	}
	return f.db.Model(&oldFood).Updates(food).Error
}

func (f *FoodRepositoryPSQL) DeleteFood(food Food) error {
	return f.db.Where("name = ?", food.Name).Delete(&food).Error
}

func (f *FoodRepositoryPSQL) CreateRecord(r Record) error {
	var food Food
	if err := f.db.Where("name = ?", r.FoodName).First(&food).Error; err != nil {
		return err
	}
	return f.db.Model(&food).Association("Records").Append(&r)
}

func (f *FoodRepositoryPSQL) GetRecordsByDate(year, month, day int64) ([]Record, error) {
	var results []Record
	var startTime, endTime time.Time
	if day > 0 {
		startTime = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
		endTime = time.Date(int(year), time.Month(month), int(day+1), 0, 0, 0, 0, time.Local).Add(-time.Second)
	} else {
		startTime = time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(int(year), time.Month(month+1), 1, 0, 0, 0, 0, time.Local).Add(-time.Second)
	}
	f.db.Where("eating_date BETWEEN ? AND ?", startTime, endTime).Order("updated_at").Find(&results)
	return results, nil
}

func (f *FoodRepositoryPSQL) UpdateRecord(oldRecord, record Record) error {
	var food Food
	if err := f.db.Where("name = ?", oldRecord.FoodName).First(&food).Error; err != nil {
		return err
	}
	return f.db.Model(&food).Association("Records").Replace(&oldRecord, &record)
}

func (f *FoodRepositoryPSQL) DeleteRecord(record Record) error {
	return f.db.Where("eating_date = ?", record.EatingDate).Delete(&record).Error
}
