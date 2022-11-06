package controller

import (
	"encoding/json"
	"foodie_manager/db"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecords(repo db.FoodRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []byte
		var err error
		var record db.Record

		if data, err = ioutil.ReadAll(c.Request.Body); err != nil {
			c.JSON(400, gin.H{"message": "Error when reading request body."})
			return
		}
		if err = json.Unmarshal(data, &record); err != nil {
			c.JSON(400, gin.H{"message": "Error when parsing request body."})
			return
		}

		if err = repo.CreateRecord(record); err != nil {
			c.JSON(500, gin.H{"message": "Something error when creating record."})
			return
		}
		c.JSON(201, gin.H{
			"message": "Record created",
		})
	}
}

func GetRecordsByDate(repo db.FoodRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		year, month, shouldReturn := parseDate(c)
		if shouldReturn {
			return
		}
		day, err := strconv.ParseInt(c.Param("day"), 10, 8)
		if err != nil {
			day = 0
		}
		if day < 0 || day > 31 {
			c.JSON(400, gin.H{"message": "Day should be within 0-31"})
			return
		}

		records, err := repo.GetRecordsByDate(year, month, day)
		if err != nil {
			c.JSON(500, gin.H{"message": "Service unavailable."})
			return
		}
		c.JSON(200, records)
	}
}

func UpdateRecordsByDate(repo db.FoodRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		year, month, shouldReturn := parseDate(c)
		if shouldReturn {
			return
		}

		day, err := strconv.ParseInt(c.Param("day"), 10, 8)
		if err != nil {
			c.JSON(400, gin.H{"message": "Cannot parse day"})
			return
		}
		if day < 0 || day > 31 {
			c.JSON(400, gin.H{"message": "Day should be within 0-31"})
			return
		}

		// Parse new record in request body
		var data []byte
		var record db.Record
		if data, err = ioutil.ReadAll(c.Request.Body); err != nil {
			c.JSON(400, gin.H{"message": "Error when reading request body."})
			return
		}
		if err = json.Unmarshal(data, &record); err != nil {
			c.JSON(400, gin.H{"message": "Error when parsing request body."})
			return
		}

		oldRecord, err := repo.GetRecordsByDate(year, month, day)
		if len(oldRecord) == 0 || err != nil {
			c.JSON(400, gin.H{"message": "Record not found."})
			return
		}
		if err = repo.UpdateRecord(oldRecord[0], record); err != nil {
			c.JSON(500, gin.H{"message": "Service unavailable."})
			return
		}
		c.JSON(200, gin.H{"message": "Record updated."})
	}
}

func parseDate(c *gin.Context) (int64, int64, bool) {
	year, err := strconv.ParseInt(c.Param("year"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Cannot parse year"})
		return 0, 0, true
	}
	if year < 1 {
		c.JSON(400, gin.H{"message": "Year should be >= 1"})
		return 0, 0, true
	}

	month, err := strconv.ParseInt(c.Param("month"), 10, 8)
	if err != nil {
		c.JSON(400, gin.H{"message": "Cannot parse month"})
		return 0, 0, true
	}
	if month < 1 || month > 12 {
		c.JSON(400, gin.H{"message": "Month should be within 1-12"})
		return 0, 0, true
	}
	return year, month, false
}

func SetupRecordControllers(r *gin.Engine, repo db.FoodRepository) *gin.Engine {
	r.POST("/records", CreateRecords(repo))
	r.PUT("/records/:year/:month/:day", UpdateRecordsByDate(repo))
	r.GET("/records/:year/:month/:day", GetRecordsByDate(repo))
	r.GET("/records/:year/:month", GetRecordsByDate(repo))
	return r
}
