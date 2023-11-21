package controllers

import (
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func Login(ctx *gin.Context) {
	// var input LoginInput
	// user := models.Users{}
	// var err error

	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return err
	// }

	// err = initializers.DB.Where("email=?", input.Email).Take(&user).Error

	// if err != nil {
	// 	// return "", err
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

}
