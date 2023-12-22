package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/http/controllers/users_controller"
)

func UsersHandle(r *gin.Engine) {
	route := r.Group("/users")

	route.Use()
	{
		route.POST("", users_controller.Store)
		route.PATCH("/:userId", users_controller.Update)
	}
}
