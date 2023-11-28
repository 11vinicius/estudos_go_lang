package routes

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("Authorization")
		claims := jwt.MapClaims{}

		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Não autorization"),
			})
			ctx.Abort()
			return
		}

		_, err := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(clientToken), nil
		})

		if err != nil {
			ctx.Next()
			return

		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Token expirado"),
		})

	}
}
