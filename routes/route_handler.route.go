package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteHandler struct{}

func NewRoutehandler() *RouteHandler {
	return &RouteHandler{}
}

func (handler *RouteHandler) Handle() {
	router := handler.GetRouteEngine()

	router.Run() // run server
}

func (handler *RouteHandler) GetRouteEngine() *gin.Engine {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "API IS ON FIRE!"})
	})

	NewUsersRoutes().Handle(router)

	return router
}
