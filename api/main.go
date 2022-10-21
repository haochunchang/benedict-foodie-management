package main

import (
	"fmt"
	"foodie_manager/db"
	"os"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hell World",
	})
}

func main() {
	connectionInfo := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=unittest port=5432 sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("username"), os.Getenv("password"),
	)
	conn := db.GetDB(connectionInfo)

	r := gin.Default()
	r.GET("/", HelloWorld)
	r.Run() // listen and serve on 0.0.0.0:8080
}
