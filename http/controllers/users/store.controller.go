package users

import (
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/actions/users_action"
	"github.com/joeltiago00/first-api-go-lang/http/requests/users_request"
	"net/http"
)

type StoreController struct{}

func NewStoreController() *StoreController {
	return &StoreController{}
}

func (r *StoreController) Store(context *gin.Context) {
	var payload map[string]string

	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	hasError, err := users_request.ValidateStore(payload)

	if hasError {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": users_action.Store(payload)})
}
