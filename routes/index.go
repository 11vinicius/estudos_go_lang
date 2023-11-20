package routes

import (
	"github.com/gin-gonic/gin"
)

func Index() {

	router := gin.Default()
	UserRoutes(router)

	router.Run(":8000")
}
