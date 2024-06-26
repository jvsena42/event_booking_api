package routes

import (
	"log"
	"net/http"
	"strconv"

	"com.go/event_booking/model"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		log.Println("ERROR getEvents: ", err)
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		log.Println("ERROR createEvent: ", err)
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	errorSave := event.Save()

	if errorSave != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		log.Println("ERROR createEvent: ", errorSave)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID"})
		log.Println("ERROR getEvent: ", err)
		return
	}

	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println("ERROR getEvent: ", err)
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID"})
		log.Println("ERROR UpdateEvent: ", err)
		return
	}

	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println("ERROR UpdateEvent: ", err)
		return
	}

	userId := context.GetInt64("userId")

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent model.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		log.Println("ERROR UpdateEvent: ", err)
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		log.Println("ERROR UpdateEvent: ", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID"})
		log.Println("ERROR deleteEvent: ", err)
		return
	}

	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println("ERROR deleteEvent: ", err)
		return
	}

	userId := context.GetInt64("userId")

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		log.Println("ERROR deleteEvent: ", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
