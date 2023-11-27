package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index() {

	router := gin.Default()
	UserRoutes(router)
	authRoute(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"run": "server is running"})
	})

	router.Run(":3000")
}
