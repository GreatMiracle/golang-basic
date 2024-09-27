package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/controller"
	"rest-api/db"
)

func main() {
	db.InitDB()

	// Khởi tạo một router mới
	router := gin.Default()

	router.GET("/events", controller.GetEvents)
	router.POST("/events", controller.CreateEvent)

	err := router.Run(":8082")
	if err != nil {
		return
	}

}
