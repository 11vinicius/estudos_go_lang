package routes

import (
	"boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")

	user.POST("/", controllers.UserCreate)
	user.GET("/", AuthMiddleware(), controllers.Users)
	user.GET("/:email", controllers.FindUserWhereEmail)
}
