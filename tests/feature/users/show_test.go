package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/helpers/db"
	"github.com/joeltiago00/first-api-go-lang/http/controllers/users"
	"github.com/joeltiago00/first-api-go-lang/http/resources"
	"github.com/joeltiago00/first-api-go-lang/infrastructure/factories"
	"github.com/joeltiago00/first-api-go-lang/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	route := routes.Group("/users")
	route.GET("/:userId", users.NewShowController().Show)

	return routes
}

func DeleteUser(userId int) {
	//TODO:: handle with exclusion data
	db.Database().Exec(fmt.Sprintf("delete from users where id = %d", userId))
}

func TestSuccess(test *testing.T) {
	tests.FeatureSetup(test)

	user := factories.NewUserFactory().Create()
	defer DeleteUser(user.ID)

	route := setupRoutes()

	request, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	var userResponse resources.UserResource
	json.Unmarshal(response.Body.Bytes(), &userResponse)

	assert.Equal(test, user.FirstName, userResponse.Data.FirstName)
	assert.Equal(test, http.StatusOK, response.Code)
}

func TestUserNotExists(test *testing.T) {
	tests.FeatureSetup(test)

	route := setupRoutes()

	request, _ := http.NewRequest("GET", "/users/100000", nil)
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	assert.Equal(test, `{"errors":"User 100000 not found."}`, string(response.Body.Bytes()))
	assert.Equal(test, http.StatusNotFound, response.Code)
}
