package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joeltiago00/first-api-go-lang/http/resources"
	"github.com/joeltiago00/first-api-go-lang/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type storeUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func TestStoreUserSuccess(test *testing.T) {
	tests.FeatureSetup(test)

	route := tests.SetupRoutes()

	user := storeUserPayload{FirstName: "Test", LastName: "Testing", Email: "test@test.com"}
	payload, _ := json.Marshal(user)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	var userResponse resources.UserResource
	json.Unmarshal(response.Body.Bytes(), &userResponse)
	defer tests.DeleteUser(userResponse.Data.ID)

	today := time.Now()

	assert.True(test, userResponse.Data.ID != 0)
	assert.Equal(
		test,
		fmt.Sprintf(
			"%d-%d-%d",
			userResponse.Data.CreatedAt.Year(),
			userResponse.Data.CreatedAt.Month(),
			userResponse.Data.CreatedAt.Day(),
		),
		fmt.Sprintf(
			"%d-%d-%d",
			today.Year(),
			today.Month(),
			today.Day(),
		),
	)
}

func TestStoreUserFailInvalidFirstName(test *testing.T) {
	tests.FeatureSetup(test)

	route := tests.SetupRoutes()

	user := storeUserPayload{FirstName: "ab", LastName: "Testing", Email: "test@test.com"}
	payload, _ := json.Marshal(user)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	assert.Equal(test, response.Code, http.StatusUnprocessableEntity)
	assert.Equal(
		test,
		`{"errors":{"first_name":{"0":"Field first_name must be a least 3 characters."}}}`,
		string(response.Body.Bytes()),
	)
}

func TestStoreUserFailInvalidLastName(test *testing.T) {
	tests.FeatureSetup(test)

	route := tests.SetupRoutes()

	user := storeUserPayload{FirstName: "abc", LastName: "ab", Email: "test@test.com"}
	payload, _ := json.Marshal(user)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	assert.Equal(test, response.Code, http.StatusUnprocessableEntity)
	assert.Equal(
		test,
		`{"errors":{"last_name":{"0":"Field last_name must be a least 3 characters."}}}`,
		string(response.Body.Bytes()),
	)
}

func TestStoreUserFailInvalidEmail(test *testing.T) {
	tests.FeatureSetup(test)

	route := tests.SetupRoutes()

	user := storeUserPayload{FirstName: "abc", LastName: "abc", Email: "test"}
	payload, _ := json.Marshal(user)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	assert.Equal(test, response.Code, http.StatusUnprocessableEntity)
	assert.Equal(
		test,
		`{"errors":{"email":{"0":"Field must be a valid email."}}}`,
		string(response.Body.Bytes()),
	)
}

func TestStoreUserFailInvalidPayload(test *testing.T) {
	tests.FeatureSetup(test)

	route := tests.SetupRoutes()

	user := storeUserPayload{FirstName: "ab", LastName: "ab", Email: "test"}
	payload, _ := json.Marshal(user)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()

	route.ServeHTTP(response, request)

	assert.Equal(test, response.Code, http.StatusUnprocessableEntity)
	assert.Equal(
		test,
		`{"errors":{"email":{"0":"Field must be a valid email."},"first_name":{"0":"Field first_name must be a least 3 characters."},"last_name":{"0":"Field last_name must be a least 3 characters."}}}`,
		string(response.Body.Bytes()),
	)
}
