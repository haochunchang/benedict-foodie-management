package controller

import (
	"encoding/json"
	"fmt"
	"foodie_manager/db"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecords(repo db.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []byte
		var err error
		var record []db.Record

		if data, err = ioutil.ReadAll(c.Request.Body); err != nil {
			c.JSON(400, gin.H{"message": "Error when reading request body."})
			return
		}
		if err = json.Unmarshal(data, &record); err != nil {
			c.JSON(400, gin.H{"message": "Error when parsing request body."})
			return
		}

		for _, r := range record {
			if err = repo.CreateRecord(r); err != nil {
				c.JSON(500, gin.H{"message": "Something error when creating record."})
				return
			}
		}
		c.JSON(201, gin.H{
			"message": "Record created",
		})
	}
}

func GetRecordsByDate(repo db.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		year, err := strconv.ParseInt(c.Param("year"), 10, 64)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"message": "Cannot parse year"})
			return
		}
		if year < 1 {
			c.JSON(400, gin.H{"message": "Year should be >= 1"})
			return
		}

		month, err := strconv.ParseInt(c.Param("month"), 10, 8)
		if err != nil {
			c.JSON(400, gin.H{"message": "Cannot parse month"})
			return
		}
		if month < 1 || month > 12 {
			c.JSON(400, gin.H{"message": "Month should be within 1-12"})
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
		if len(records) == 0 {
			c.JSON(400, gin.H{"message": "Records not found"})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"message": "Service unavailable."})
			return
		}
		c.JSON(200, records)
	}
}

func SetupRecordControllers(r *gin.Engine, record db.RecordRepository) *gin.Engine {
	r.POST("/records", CreateRecords(record))
	r.GET("/records/:year/:month/:day", GetRecordsByDate(record))
	r.GET("/records/:year/:month", GetRecordsByDate(record))
	return r
}
