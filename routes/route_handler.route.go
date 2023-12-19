package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() {
	r := gin.Default()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "API IS ON FIRE!"})
	})

	UsersHandle(r)

	r.Run() // run server
}
