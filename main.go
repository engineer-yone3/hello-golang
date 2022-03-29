package main

import (
	"rest-api-test/entity"
	"rest-api-test/server"
	// "rest-api-test/controller"
	"net/http"
	"rest-api-test/database"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DB
// var err error


func dbInit() {
	db := database.GormConnect()

	db.AutoMigrate(&entity.SampleEntity{})

	defer db.Close()
}

func main() {
	// r := server.GetRouter()
	// r.Run(":8080")

	dbInit()
	
	server.GetRouterHandle()
	http.ListenAndServe(":8080", nil)
}