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
		Name:        "hororo",
		Type:        "wet",
		Description: "A kind of can food",
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

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foods/hororo", nil)
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
	if resp.Description != food.Description {
		t.Errorf("Food Description not correct\nExpected %s, got %s", food.Description, resp.Description)
	}
}

func TestUpdateFoodByNameRoute(t *testing.T) {
	router := SetupFoodControllers(gin.Default(), repo)
	food := getSampleFood()
	food.Name = "baily"
	json_data, _ := json.Marshal(food)
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	// PUT should respond success
	food.Name = "hihi"
	json_data, _ = json.Marshal(food)

	w := httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/foods?oldFoodName=baily", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d", w.Code)
	}
	if w.Body.String() != fmt.Sprintf(`{"message":"Food %s updated to %s"}`, "baily", "hihi") {
		t.Errorf("Food name not updated, got %s", w.Body.String())
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/foods/baily", nil)
	router.ServeHTTP(w, req)
	if w.Code == 200 {
		t.Errorf("Get by old name should not be found, got %s", w.Body.String())
	}

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
