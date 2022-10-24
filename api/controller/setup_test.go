package controller

import (
	"fmt"
	"foodie_manager/db"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testingConnectionInfo string = fmt.Sprintf(
	"host=localhost user=%s password=%s dbname=unittest port=5432 sslmode=disable TimeZone=Asia/Taipei",
	os.Getenv("username"), os.Getenv("password"),
)

var conn *gorm.DB = db.GetDB(testingConnectionInfo)
var foodRepo *db.FoodRepositoryPSQL = db.NewFoodRepositoryPSQL(conn)
var recordRepo *db.RecordRepositoryPSQL = db.NewRecordRepositoryPSQL(conn)

func TestMain(m *testing.M) {
	recordRepo.Clear()
	foodRepo.Clear()

	foodRepo.Init()
	recordRepo.Init()

	code := m.Run()

	recordRepo.Clear()
	foodRepo.Clear()
	os.Exit(code)
}
