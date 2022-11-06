package controller

import (
	"fmt"
	"foodie_manager/db"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var testingConnectionInfo string = fmt.Sprintf(
	"host=localhost user=%s password=%s dbname=unittest port=5432 sslmode=disable TimeZone=Asia/Taipei",
	os.Getenv("username"), os.Getenv("password"),
)

var conn *gorm.DB = db.GetDB(testingConnectionInfo)
var repo *db.FoodRepositoryPSQL = db.NewFoodRepositoryPSQL(conn)

func TestMain(m *testing.M) {
	repo.Clear()
	repo.Init()
	gin.SetMode(gin.ReleaseMode)

	code := m.Run()

	repo.Clear()
	os.Exit(code)
}
