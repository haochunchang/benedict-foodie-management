package db

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	repo.Clear()
	repo.Init()

	date := time.Date(2022, 10, 21, 0, 0, 1, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date, CurrentQuantity: 24}

	target := Record{
		Food:       food,
		EatingDate: date,
	}
	if err := repo.CreateRecord(target); err != nil {
		panic("Failed to create record: " + err.Error())
	}
	code := m.Run()

	repo.Clear()
	os.Exit(code)
}

func TestCreateFood(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}
	if err := repo.CreateFood(food); err != nil {
		t.Errorf("Failed to create food, got %v", err)
	}
}

func TestGetFood(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 1, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}

	res := repo.GetFoodByName(food.Name)
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

func TestUpdateFoodByName(t *testing.T) {
	name := "hororo"
	newName := "hihi"
	newFood := Food{Name: newName}
	if err := repo.UpdateFoodByName(name, newFood); err != nil {
		t.Errorf("Failed to update food, got %v", err)
	}
	res := repo.GetFoodByName(newName)
	if res.Name != newName {
		t.Errorf("Incorrect food name\nExpect %s, got %s", newName, res.Name)
	}
}

func TestDeleteFood(t *testing.T) {
	name := "hororo"
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: name, Type: "wet", PurchaseDate: date}
	if repo.CreateFood(food) != nil {
		t.Fatal("Failed to create food")
	}
	if err := repo.DeleteFood(food); err != nil {
		t.Errorf("Failed to delete food, got %v", err)
	}
}

func TestGetRecordByDate(t *testing.T) {
	date := time.Date(2022, 10, 31, 0, 0, 1, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo2", Type: "wet", PurchaseDate: date}

	target := Record{
		Food:       food,
		EatingDate: date,
	}
	if err := repo.CreateRecord(target); err != nil {
		t.Fatalf("Failed to create record: %v", err.Error())
	}

	result, err := repo.GetRecordsByDate(2022, 10, 31)
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

	result, err = repo.GetRecordsByDate(2022, 10, 0)
	if err != nil {
		t.Errorf("Failed to get record by month, got %v", err)
	}
	if len(result) != 2 || result[1].EatingDate != date {
		t.Errorf("Eating date is incorrect\n Expect %s, got %v", date, result[1].EatingDate)
	}
}

func TestDeleteRecord(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}

	target := Record{
		Food:       food,
		EatingDate: date,
	}
	if repo.CreateRecord(target) != nil {
		t.Fatal("Failed to create record")
	}
	if err := repo.DeleteRecord(target); err != nil {
		t.Errorf("Failed to delete record, got %v", err)
	}
}

func TestUpdateRecordByDate(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	food := Food{Name: "hororo", Type: "wet", PurchaseDate: date}

	target := Record{
		Food:       food,
		EatingDate: date,
	}
	newRecord := target
	newRecord.EatenQuantity = 10
	if err := repo.UpdateRecordByDate(2022, 10, 21, newRecord); err != nil {
		t.Errorf("Failed to update record, got %v", err)
	}

	resp, err := repo.GetRecordsByDate(2022, 10, 21)
	if err != nil {
		t.Errorf("Failed to get record, got %v", err)
	}
	if len(resp) != 1 || resp[0].EatenQuantity != 10 {
		t.Errorf("Incorrect updated record, got %v", resp)
	}
}
