package routes

import (
	"boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func authRoute(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.POST("/signin", controllers.Login)

}
