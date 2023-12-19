package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersHandle(r *gin.Engine) {
	route := r.Group("/users")

	route.Use()
	{
		route.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "API IS ON FIRE!"})
		})
	}
}
