package routes

import "github.com/gin-gonic/gin"

func authRoute(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.Post("/login", controllers.login)

}
