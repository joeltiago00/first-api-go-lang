package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/http/controllers/users"
)

func UsersHandle(r *gin.Engine) {
	route := r.Group("/users")

	route.Use()
	{
		route.POST("", users.NewStoreController().Store)
		route.PATCH("/:userId", users.NewUpdateController().Update)
		route.GET("/:userId", users.NewShowController().Show)
	}
}
