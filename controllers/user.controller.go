package controllers

import (
	"net/http"

	"github.com/JulesAD96/go-jwt-auth/database"
	"github.com/JulesAD96/go-jwt-auth/models"
	"github.com/gin-gonic/gin"
)

// Register a new user by using credentials
func Register(context *gin.Context) {
	var user models.User

	// binding data
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// hash password
	errPasswordHashing := user.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": errPasswordHashing.Error()})
		context.Abort()
		return
	}

	//create record in database
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	// response
	context.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}
