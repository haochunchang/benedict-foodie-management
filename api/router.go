package main

import (
	"foodie_manager/controller"
	"foodie_manager/db"
	docs "foodie_manager/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is the backend of Benedict Foodie Management",
	})
}

func setupRouter(repos map[string]interface{}) *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/", HelloWorld)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	food, ok := repos["food"].(*db.FoodRepositoryPSQL)
	if !ok {
		panic("Food repository is not configured")
	}
	food.Init()
	controller.SetupFoodControllers(r, food)

	record, ok := repos["record"].(*db.RecordRepositoryPSQL)
	if !ok {
		panic("Record repository is not configured")
	}
	record.Init()
	controller.SetupRecordControllers(r, record)

	return r
}
