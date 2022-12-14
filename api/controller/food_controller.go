package controller

import (
	"encoding/json"
	"fmt"
	"foodie_manager/db"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

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

func UpdateFoodByName(repo db.FoodRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("oldFoodName")

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
		if err = repo.UpdateFoodByName(name, food); err != nil {
			c.JSON(500, gin.H{"message": "Something error when updating food."})
			return
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Food %s updated to %s", name, food.Name),
		})
	}
}

func SetupFoodControllers(r *gin.Engine, repo db.FoodRepository) *gin.Engine {
	r.GET("/foods/:name", GetFoodByName(repo))
	r.POST("/foods", CreateFood(repo))
	r.PUT("/foods", UpdateFoodByName(repo))
	return r
}
