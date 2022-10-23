package main

import (
	"encoding/json"
	"fmt"
	"foodie_manager/db"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hell World",
	})
}

func CreateFood(repo db.FoodRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []byte
		var err error
		var food db.Food

		if data, err = ioutil.ReadAll(c.Request.Body); err != nil {
			c.JSON(400, gin.H{"message": "Error when reading request body."})
			return
		}

		if err = json.Unmarshal(data, &food); err != nil {
			c.JSON(400, gin.H{"message": "Error when parsing request body."})
			return
		}

		if err = repo.CreateFood(food); err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Something error when creating food."})
			return
		}
		c.JSON(201, gin.H{
			"message": "Food created",
		})
	}
}

func GetFoodByName(repo db.FoodRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		food := repo.GetFoodByName(name)
		if len(food.PurchaseDate) == 0 {
			c.JSON(400, gin.H{"message": "Food not found."})
			return
		}
		c.JSON(200, food)
	}
}

func setupRouter(repos map[string]interface{}) *gin.Engine {
	food, ok := repos["food"].(*db.FoodRepositoryPSQL)
	if !ok {
		panic("Food repository is not configured")
	}
	// record, ok := repos["record"].(*db.RecordRepositoryPSQL)
	// if !ok {
	// 	panic("Record repository is not configured")
	// }
	r := gin.Default()
	r.GET("/", HelloWorld)
	r.POST("/foods", CreateFood(food))
	r.GET("/foods/:name", GetFoodByName(food))
	return r
}
