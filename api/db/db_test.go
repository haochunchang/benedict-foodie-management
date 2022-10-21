package db

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

var testingConnectionInfo string = fmt.Sprintf(
	"host=localhost user=%s password=%s dbname=unittest port=5432 sslmode=disable TimeZone=Asia/Taipei",
	os.Getenv("username"), os.Getenv("password"),
)

var Db *gorm.DB = GetDB(testingConnectionInfo)

func TestMain(m *testing.M) {
	clearTables()
	InitDB(testingConnectionInfo)
	code := m.Run()
	clearTables()
	os.Exit(code)
}

func clearTables() {
	Db.Exec("DROP TABLE IF EXISTS records")
	Db.Exec("DROP TABLE IF EXISTS foods")
}

func TestCreateFood(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}
	if err := CreateFood(Db, food); err != nil {
		t.Errorf("Failed to create food, got %v", err)
	}
}

func TestGetFood(t *testing.T) {
	name := "hororo"
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: name, Type: "wet", PurchaseDate: date}
	if CreateFood(Db, food) != nil {
		t.Fatal("Failed to create food")
	}

	res := GetFoodByName(Db, name)
	if res.Name != food.Name {
		t.Errorf("Incorrect food name\nExpect %v, got %v", food.Name, res.Name)
	}
	if res.Type != food.Type {
		t.Errorf("Incorrect food type\nExpect %v, got %v", food.Type, res.Type)
	}
	if res.PurchaseDate != food.PurchaseDate {
		t.Errorf("Incorrect food purchase date\nExpect %v, got %v", food.PurchaseDate, res.PurchaseDate)
	}
}

func TestGetRecordByDate(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}
	if CreateFood(Db, food) != nil {
		t.Fatal("Failed to create food")
	}

	target := Record{
		Food:       food,
		EatingDate: date,
	}
	if CreateRecord(Db, target) != nil {
		t.Fatal("Failed to create record")
	}
	result, err := GetRecordsByDate(Db, date)
	if err != nil {
		t.Errorf("Failed to get record by date, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Number of record is incorrect, got %d", len(result))
	}
	if result[0].EatingDate != date {
		t.Errorf("Eating date is incorrect\n Expect %s, got %v", date, result[0].EatingDate)
	}
	if result[0].Food.Name != food.Name {
		t.Errorf("Incorrect food\n Expect %v, got %v", food.Name, result[0].Food.Name)
	}
}
