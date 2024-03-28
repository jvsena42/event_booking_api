package main

import (
	"fmt"
	"net/http"

	"com.go/event_booking/db"
	"com.go/event_booking/model"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()

	if err != nil {
		fmt.Println("Error getEvents: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.Id = 1
	event.UserId = 1

	errorSave := event.Save()

	if errorSave != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
