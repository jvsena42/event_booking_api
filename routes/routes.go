package routes

import (
	"com.go/event_booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticaded := server.Group("/")
	authenticaded.Use(middlewares.Autenticate)
	authenticaded.POST("/events", createEvent)
	authenticaded.PUT("/events/:id", updateEvent)
	authenticaded.DELETE("events/:id", deleteEvent)
	authenticaded.POST("events/:id/register", registerForEvent)
	authenticaded.DELETE("events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
