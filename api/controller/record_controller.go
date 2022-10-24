package controller

import (
	"encoding/json"
	"fmt"
	"foodie_manager/db"
	"io/ioutil"
	"time"

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
				fmt.Println(err)
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
		date := c.Param("date")

		if _, err := time.Parse(time.RFC3339, date); err != nil {
			c.JSON(400, gin.H{"message": "Date time format needs to be RFC3339"})
			return
		}

		records, err := repo.GetRecordsByDate(date)
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
	r.GET("/records/:date", GetRecordsByDate(record))
	r.POST("/records", CreateRecords(record))
	return r
}
