package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/http/controllers/users"
)

type UsersRoutes struct{}

func NewUsersRoutes() *UsersRoutes {
	return &UsersRoutes{}
}

func (r *UsersRoutes) Handle(router *gin.Engine) {
	route := router.Group("/users")

	route.Use()
	{
		route.POST("", users.NewStoreController().Store)
		route.PATCH("/:userId", users.NewUpdateController().Update)
		route.GET("/:userId", users.NewShowController().Show)
	}
}
