package routes

import (
	"log"
	"net/http"

	"com.go/event_booking/model"
	"com.go/event_booking/utils"
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

func login(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		log.Println("ERROR login: ", err)
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		log.Println("ERROR login: ", err)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("ERROR login: ", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token, "user": user})
}
