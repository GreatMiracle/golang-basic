package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/service"
)

//func GetEvents(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
//}

func GetEvents(context *gin.Context) {
	events, _ := service.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	var req models.Event
	err := context.ShouldBindJSON(&req)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	req.ID = 1
	req.UserID = 1

	if err := service.SaveEvent(&req); err != nil {
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": req})

}
