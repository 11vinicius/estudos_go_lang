package controllers

import (
	"boilerplate/initializers"
	"boilerplate/models"
	"boilerplate/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var input LoginInput
	user := models.Users{}
	var err error

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = initializers.DB.Where("email = ?", input.Email).Find(&user).Error

	if err != nil {
		return
	}

	err = utils.VerifyPassword(input.Password, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email e/ou senha inválida"})
		return

	}
	token, err := createToken(user)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func createToken(user models.Users) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	return token.SignedString([]byte("hmacSampleSecret"))
}
