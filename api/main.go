package main

import (
	"github.com/gin-gonic/gin"
)

var connectionInfo = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Taipei"

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hell World",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", HelloWorld)

	r.GET("/")
	r.Run() // listen and serve on 0.0.0.0:8080
}
