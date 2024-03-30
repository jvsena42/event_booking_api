package routes

import (
	"log"
	"net/http"
	"strconv"

	"com.go/event_booking/model"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID"})
		log.Println("ERROR Register event: ", err)
		return
	}

	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println("ERROR Register event: ", err)
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		log.Println("ERROR Register event: ", err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {

}
