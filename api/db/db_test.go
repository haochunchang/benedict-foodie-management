package db

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	foodRepo.Clear()
	recordRepo.Clear()

	foodRepo.Init()
	recordRepo.Init()

	code := m.Run()

	foodRepo.Clear()
	recordRepo.Clear()
	os.Exit(code)
}

func TestCreateFood(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}
	if err := foodRepo.CreateFood(food); err != nil {
		t.Errorf("Failed to create food, got %v", err)
	}
}

func TestGetFood(t *testing.T) {
	name := "hororo"
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: name, Type: "wet", PurchaseDate: date}
	if foodRepo.CreateFood(food) != nil {
		t.Fatal("Failed to create food")
	}

	res := foodRepo.GetFoodByName(name)
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
	if foodRepo.CreateFood(food) != nil {
		t.Fatal("Failed to create food")
	}

	target := Record{
		FoodName:   food.Name,
		EatingDate: date,
	}
	if recordRepo.CreateRecord(target) != nil {
		t.Fatal("Failed to create record")
	}
	result, err := recordRepo.GetRecordsByDate(date)
	if err != nil {
		t.Errorf("Failed to get record by date, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Number of record is incorrect, got %d", len(result))
	}
	if result[0].EatingDate != date {
		t.Errorf("Eating date is incorrect\n Expect %s, got %v", date, result[0].EatingDate)
	}
	if result[0].FoodName != food.Name {
		t.Errorf("Incorrect food\n Expect %v, got %v", food.Name, result[0].FoodName)
	}
}
