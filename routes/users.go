package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse body"})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create the user. Try again later!"})
		return
	}

	fmt.Printf(user.Email)
	fmt.Printf(user.Password)

	ctx.JSON(http.StatusOK, gin.H{"message": "user created successfully!"})
}

func loginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "invalid request"})
		return
	}

	err = user.VerifyCredentials()

	if err != nil {
		fmt.Printf(`%v\n`, err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful"})

}
