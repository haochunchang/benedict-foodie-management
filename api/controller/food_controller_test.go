package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"foodie_manager/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

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

func TestCreateFoodRoute(t *testing.T) {
	router := SetupFoodControllers(gin.Default(), repo)

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
	router := SetupFoodControllers(gin.Default(), repo)

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

func TestUpdateFoodByNameRoute(t *testing.T) {
	router := SetupFoodControllers(gin.Default(), repo)

	// POST sample food
	food := getSampleFood()
	json_data, _ := json.Marshal(food)
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	// PUT should respond success
	food.Name = "hihi"
	json_data, _ = json.Marshal(food)

	w := httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/foods?oldFoodName=hororo", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d", w.Code)
	}
	if w.Body.String() != fmt.Sprintf(`{"message":"Food %s updated to %s"}`, "hororo", "hihi") {
		t.Errorf("Food name not updated, got %s", w.Body.String())
	}

	// w = httptest.NewRecorder()
	// req, _ = http.NewRequest("GET", "/foods/hororo", nil)
	// router.ServeHTTP(w, req)
	// if w.Code != 200 || w.Body.String() != `"message": "Food not found."` {
	// 	t.Errorf("Get by old name should not be found, got %s", w.Body.String())
	// }

	// Get new name should get data
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/foods/hihi", nil)
	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d", w.Code)
	}

	var resp db.Food
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Decode response error: %v", err)
	}
	fmt.Print(resp)
	if resp.Name != "hihi" {
		t.Errorf("Food name not correct\nExpected hihi, got %s", resp.Name)
	}
}

func TestDeleteFoodRoute(t *testing.T) {

}
