package main

import (
	"net/http"

	"com.go/event_booking/model"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := model.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
