package routes

import (
	"log"
	"net/http"

	"com.go/event_booking/model"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		log.Println("ERROR signup: ", err)
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		log.Println("ERROR signup: ", err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
