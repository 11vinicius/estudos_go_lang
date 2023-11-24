package controllers

import (
	"boilerplate/initializers"
	"boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
}

func verify(password, passwordHashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}

// func createToken(user models.Users) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"userId": user.Id,
// 		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
// 	})

// 	// tokenString, err := token.SignedString(hmacSampleSecret)
// 	fmt.Println(tokenString, err)

// }
