package routes

import (
	"github.com/gin-gonic/gin"
)

func Index() {

	router := gin.Default()
	UserRoutes(router)

	authRoute(router)

	router.Run(":3000")
}
