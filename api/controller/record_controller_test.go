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

func TestCreateRecordRoute(t *testing.T) {
	router := SetupRecordControllers(gin.Default(), recordRepo)

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
	router := SetupRecordControllers(gin.Default(), recordRepo)

	// Add food
	food := getSampleFood()
	json_data, _ := json.Marshal(food)
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	// Add records
	records := getSampleRecords()
	json_data, _ = json.Marshal(records)
	req, _ = http.NewRequest("POST", "/records", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	// Test get record
	w := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/records/%d/%d/%d", 2022, 10, 23), nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp []db.Record
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp) == 0 {
		t.Errorf("Response not correct, got %d", len(resp))
	}
}

func TestGetRecordsByMonthRoute(t *testing.T) {
	router := SetupRecordControllers(gin.Default(), recordRepo)

	// Add food
	food := getSampleFood()
	json_data, _ := json.Marshal(food)
	req, _ := http.NewRequest("POST", "/foods", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	// Add records
	records := getSampleRecords()
	json_data, _ = json.Marshal(records)
	req, _ = http.NewRequest("POST", "/records", bytes.NewBuffer(json_data))
	router.ServeHTTP(httptest.NewRecorder(), req)

	// Test get record
	w := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/records/%d/%d", 2022, 10), nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code not 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp []db.Record
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp) == 0 {
		t.Errorf("Response not correct, got %d", len(resp))
	}
}
