package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	events, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user."})
		return
	}

	context.JSON(http.StatusOK, user)
}

func createUser(context *gin.Context) {
	var payload models.User
	err := context.ShouldBindJSON(&payload)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	user, err := payload.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created!", "user": user})
}

func updateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	_, err = models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user."})
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedUser.ID = userId
	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user."})
		return
	}

	err = user.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
