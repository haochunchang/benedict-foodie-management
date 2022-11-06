package db

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	repo.Clear()
	repo.Init()

	food := Food{Name: "hororo", Type: "wet"}
	if err := repo.CreateFood(food); err != nil {
		panic("Failed to create record: " + err.Error())
	}
	code := m.Run()

	repo.Clear()
	os.Exit(code)
}

func TestCreateFood(t *testing.T) {
	food := Food{Name: "baily", Type: "dry"}
	if err := repo.CreateFood(food); err != nil {
		t.Errorf("Failed to create food, got %v", err)
	}
}

func TestGetFood(t *testing.T) {
	food := Food{Name: "hororo", Type: "wet"}

	res, err := repo.GetFoodByName(food.Name)
	if err != nil {
		t.Errorf("Error occur when get food: %v", err)
	}
	if res.Name != food.Name {
		t.Errorf("Incorrect food name\nExpect %v, got %v", food.Name, res.Name)
	}
	if res.Type != food.Type {
		t.Errorf("Incorrect food type\nExpect %v, got %v", food.Type, res.Type)
	}
}

func TestGetRecordByDate(t *testing.T) {
	date := time.Date(2022, 10, 31, 0, 0, 1, 0, time.Local)
	food := Food{Name: "hororo", Type: "wet"}

	target := Record{
		FoodName:   food.Name,
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
	if result[0].FoodName != food.Name {
		t.Errorf("Incorrect food\n Expect %v, got %v", food.Name, result[0].FoodName)
	}

	result, err = repo.GetRecordsByDate(2022, 10, 0)
	if err != nil {
		t.Errorf("Failed to get record by month, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Number of record is incorrect, got %d", len(result))
	}
	if result[0].EatingDate != date {
		t.Errorf("Eating date is incorrect\n Expect %s, got %v", date, result[0].EatingDate)
	}
}

func TestUpdateFoodByName(t *testing.T) {
	name := "baily"
	newName := "hihi"
	newFood := Food{Name: newName}
	if err := repo.UpdateFoodByName(name, newFood); err != nil {
		t.Errorf("Failed to update food, got %v", err)
	}
	res, err := repo.GetFoodByName(newName)
	if err != nil {
		t.Errorf("Error occur when get food: %v", err)
	}
	if res.Name != newName {
		t.Errorf("Incorrect food name\nExpect %s, got %s", newName, res.Name)
	}
}

func TestUpdateRecordByDate(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local)
	food := Food{Name: "hororo", Type: "wet"}

	target := Record{
		FoodName:   food.Name,
		EatingDate: date,
	}
	if err := repo.CreateRecord(target); err != nil {
		t.Fatalf("Failed to create record: %v", err)
	}

	newRecord := target
	newRecord.EatenQuantity = 10
	if err := repo.UpdateRecord(target, newRecord); err != nil {
		t.Errorf("Failed to update record, got %v", err)
	}

	resp, err := repo.GetRecordsByDate(2022, 10, 21)
	if err != nil {
		t.Errorf("Failed to get record, got %v", err)
	}
	if resp[len(resp)-1].EatenQuantity != 10 {
		t.Errorf("Incorrect updated record, got %v", resp)
	}
}

func TestDeleteRecord(t *testing.T) {
	date := time.Date(2022, 10, 21, 0, 0, 0, 0, time.Local)
	food := Food{Name: "hororo", Type: "wet"}

	target := Record{
		FoodName:   food.Name,
		EatingDate: date,
	}
	if repo.CreateRecord(target) != nil {
		t.Fatal("Failed to create record")
	}
	if err := repo.DeleteRecord(target); err != nil {
		t.Errorf("Failed to delete record, got %v", err)
	}
}

func TestDeleteFood(t *testing.T) {
	name := "hororo"
	food := Food{Name: name, Type: "wet"}
	if err := repo.DeleteFood(food); err != nil {
		t.Errorf("Failed to delete food, got %v", err)
	}
}
