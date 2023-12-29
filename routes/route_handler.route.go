package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() {
	router := SetupRoutes()

	router.Run() // run server
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "API IS ON FIRE!"})
	})

	NewUsersRoutes().Handle(router)

	return router
}
