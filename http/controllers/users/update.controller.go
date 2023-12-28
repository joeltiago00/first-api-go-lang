package users

import (
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/actions/users_action"
	"github.com/joeltiago00/first-api-go-lang/http/requests/users_request"
	"net/http"
	"strconv"
)

type UpdateController struct{}

func NewUpdateController() *UpdateController {
	return &UpdateController{}
}

func (r *UpdateController) Update(context *gin.Context) {
	var payload map[string]string

	userId, _ := strconv.Atoi(context.Param("userId"))

	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error1": err.Error(),
		})

		return
	}

	hasError, errors := users_request.ValidateUpdate(payload)

	if hasError {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})

		return
	}

	hasError, err := users_action.Update(userId, payload)

	if hasError {
		context.JSON(http.StatusNotFound, gin.H{"errors": err})

		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
