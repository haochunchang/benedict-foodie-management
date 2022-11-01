package db

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var testingConnectionInfo string = fmt.Sprintf(
	"host=localhost user=%s password=%s dbname=unittest port=5432 sslmode=disable TimeZone=Asia/Taipei",
	os.Getenv("username"), os.Getenv("password"),
)

var conn *gorm.DB = GetDB(testingConnectionInfo)
var repo FoodRepository = &FoodRepositoryPSQL{conn}
