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
	GetFoodByName(string) Food
	UpdateFoodByName(string, Food) error
	DeleteFood(Food) error
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

// RecordRepoistory stores the food record data
type RecordRepository interface {
	Repository
	CreateRecord(Record) error
	GetRecordsByDate(int64, int64, int64) ([]Record, error)
	UpdateRecordByDate(string, Record) error
	DeleteRecord(Record) error
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

func (rr *RecordRepositoryPSQL) GetRecordsByDate(year, month, day int64) ([]Record, error) {
	var results []Record
	var startTime, endTime time.Time
	if day > 0 {
		startTime = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
		endTime = time.Date(int(year), time.Month(month), int(day+1), 0, 0, 0, 0, time.Local).Add(-time.Second)
	} else {
		startTime = time.Date(int(year), time.Month(month), 0, 0, 0, 0, 0, time.Local)
		endTime = time.Date(int(year), time.Month(month+1), 0, 0, 0, 0, 0, time.Local).Add(-time.Second)
	}
	start := startTime.Format(time.RFC3339)
	end := endTime.Format(time.RFC3339)
	rr.db.Where("eating_date BETWEEN ? AND ?", start, end).Find(&results)
	return results, nil
}

func (rr *RecordRepositoryPSQL) UpdateRecordByDate(date string, record Record) error {
	return nil
}

func (rr *RecordRepositoryPSQL) DeleteRecord(record Record) error {
	return rr.db.Where("eating_date = ?", record.EatingDate).Delete(&record).Error
}
