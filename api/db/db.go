package db

import (
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
