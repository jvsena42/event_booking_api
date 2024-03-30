package middlewares

import (
	"net/http"

	"com.go/event_booking/utils"
	"github.com/gin-gonic/gin"
)

func Autenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not auhtorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not auhtorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
