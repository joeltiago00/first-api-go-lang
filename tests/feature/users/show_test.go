package users

import (
	"encoding/json"
	"fmt"
	"github.com/joeltiago00/first-api-go-lang/http/resources"
	"github.com/joeltiago00/first-api-go-lang/infrastructure/factories"
	"github.com/joeltiago00/first-api-go-lang/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowUserSuccess(test *testing.T) {
	tests.FeatureSetup(test)

	user := factories.NewUserFactory().Create()
	defer tests.DeleteUserById(user.ID)

	route := tests.SetupRoutes()

	request, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	var userResponse resources.UserResource
	json.Unmarshal(response.Body.Bytes(), &userResponse)

	assert.Equal(test, user.FirstName, userResponse.Data.FirstName)
	assert.Equal(test, http.StatusOK, response.Code)
}

func TestShowUserNotExists(test *testing.T) {
	tests.FeatureSetup(test)

	route := tests.SetupRoutes()

	request, _ := http.NewRequest("GET", "/users/100000", nil)
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	assert.Equal(test, `{"errors":"User 100000 not found."}`, string(response.Body.Bytes()))
	assert.Equal(test, http.StatusNotFound, response.Code)
}
