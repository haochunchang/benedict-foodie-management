package main

import (
	"foodie_manager/controller"
	"foodie_manager/db"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is the backend of Benedict Foodie Management",
	})
}

func setupRouter(repos map[string]interface{}) *gin.Engine {
	r := gin.Default()
	r.GET("/", HelloWorld)

	food, ok := repos["food"].(*db.FoodRepositoryPSQL)
	if !ok {
		panic("Food repository is not configured")
	}
	controller.SetupFoodControllers(r, food)

	record, ok := repos["record"].(*db.RecordRepositoryPSQL)
	if !ok {
		panic("Record repository is not configured")
	}
	controller.SetupRecordControllers(r, record)

	return r
}
