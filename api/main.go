package main

import (
	"fmt"
	"foodie_manager/db"
	"os"
)

func main() {
	var connectionInfo string = fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=unittest port=5432 sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("username"), os.Getenv("password"),
	)
	conn := db.GetDB(connectionInfo)

	r := setupRouter(map[string]interface{}{
		"food": db.NewFoodRepositoryPSQL(conn),
	})
	r.Run(":8080")
}
