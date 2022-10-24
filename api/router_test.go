package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"foodie_manager/db"
	"net/http"
	"net/http/httptest"
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
var repos = map[string]interface{}{
	"food":   foodRepo,
	"record": recordRepo,
}

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

func getSampleFood() db.Food {
	food := db.Food{
		Name:            "hororo",
		Type:            "wet",
		PurchaseDate:    "2022-10-22T00:00:00+08:00",
		CurrentQuantity: 24,
		Description:     "A kind of can food",
	}
	return food
}

func getSampleRecords() []db.Record {
	return []db.Record{
		{
			FoodName:          "hororo",
			Description:       "",
			EatingDate:        "2022-10-22T00:00:00+08:00",
			EatenQuantity:     0.5,
			SatisfactionScore: 3,
		},
		{
			FoodName:          "hororo",
			Description:       "",
			EatingDate:        "2022-10-23T00:00:00+08:00",
			EatenQuantity:     0.5,
			SatisfactionScore: 5,
		},
		{
			FoodName:          "hororo",
			Description:       "",
			EatingDate:        "2022-10-24T00:00:00+08:00",
			EatenQuantity:     0.5,
			SatisfactionScore: 2,
		},
	}
}

func TestCreateFoodRoute(t *testing.T) {
	router := setupRouter(repos)

	food := getSampleFood()
	json_data, _ := json.Marshal(food)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Status code not 201, got %d", w.Code)
	}
	resp := w.Body.String()
	if resp != `{"message":"Food created"}` {
		t.Errorf("Response not correct, got %s", resp)
	}
}

func TestGetFoodbyNameRoute(t *testing.T) {
	router := setupRouter(repos)

	food := getSampleFood()
	json_data, _ := json.Marshal(food)

	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	w := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/foods/hororo", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d", w.Code)
	}

	var resp db.Food
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Decode response error: %v", err)
	}
	if resp.Name != food.Name {
		t.Errorf("Food name not correct\nExpected %s, got %s", food.Name, resp.Name)
	}
	if resp.Type != food.Type {
		t.Errorf("Food Type not correct\nExpected %s, got %s", food.Type, resp.Type)
	}
	if resp.PurchaseDate != food.PurchaseDate {
		t.Errorf("Food PurchaseDate not correct\nExpected %s, got %s", food.PurchaseDate, resp.PurchaseDate)
	}
	if resp.CurrentQuantity != food.CurrentQuantity {
		t.Errorf("Food CurrentQuantity not correct\nExpected %f, got %f", food.CurrentQuantity, resp.CurrentQuantity)
	}
	if resp.Description != food.Description {
		t.Errorf("Food Description not correct\nExpected %s, got %s", food.Description, resp.Description)
	}
}

func TestCreateRecordRoute(t *testing.T) {
	router := setupRouter(repos)

	food := getSampleFood()
	json_data, _ := json.Marshal(food)
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	records := getSampleRecords()
	json_data, _ = json.Marshal(records)

	w := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/records", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Status code not 201, got %d", w.Code)
	}
	resp := w.Body.String()
	if resp != `{"message":"Record created"}` {
		t.Errorf("Response not correct, got %s", resp)
	}
}

func TestGetRecordsByDateRoute(t *testing.T) {
	router := setupRouter(repos)

	food := getSampleFood()
	json_data, _ := json.Marshal(food)
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	records := getSampleRecords()
	json_data, _ = json.Marshal(records)

	date := "2022-10-23T00:00:00+08:00"
	w := httptest.NewRecorder()
	req, _ = http.NewRequest(
		"GET",
		fmt.Sprintf("/records/%s", date),
		bytes.NewBuffer(json_data),
	)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d", w.Code)
	}

	var resp []db.Record
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp) == 0 {
		t.Errorf("Response not correct, got %d", len(resp))
	}
}
