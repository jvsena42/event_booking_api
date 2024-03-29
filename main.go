package main

import (
	"log"
	"net/http"
	"strconv"

	"com.go/event_booking/db"
	"com.go/event_booking/model"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		log.Fatal("ERROR getEvents: ", err)
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		log.Fatal("ERROR createEvent: ", err)
		return
	}

	event.Id = 1
	event.UserId = 1

	errorSave := event.Save()

	if errorSave != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		log.Fatal("ERROR createEvent: ", errorSave)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID"})
		log.Fatal("ERROR getEvent: ", err)
		return
	}

	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Fatal("ERROR getEvent: ", err)
		return
	}

	context.JSON(http.StatusOK, event)
}
