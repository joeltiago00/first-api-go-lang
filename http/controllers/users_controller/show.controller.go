package users_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/actions/users_action"
	"net/http"
	"strconv"
)

func Show(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("userId"))

	user, err := users_action.Show(userId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}
