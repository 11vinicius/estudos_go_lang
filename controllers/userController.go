package controllers

import (
	"boilerplate/initializers"
	"boilerplate/models"
	"boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCreateValidation struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func UserCreate(ctx *gin.Context) {
	var input UserCreateValidation

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.Users{Email: input.Email, Password: utils.Hashed(input.Password)}
	initializers.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func Users(ctx *gin.Context) {
	var users []models.Users

	initializers.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func FindUserWhereEmail(ctx *gin.Context) {
	var user models.Users

	email := ctx.Param("email")
	initializers.DB.Where("email=?", email).Find(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
