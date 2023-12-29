package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/helpers/db"
	"github.com/joeltiago00/first-api-go-lang/http/controllers/users"
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database"
	"testing"
)

// FeatureSetup Make application setup to run tests
func FeatureSetup(test *testing.T) {
	// Set env vars for tests
	test.Setenv("APP_ENV", "testing")

	database.Handler()
}

func DeleteUser(userId int) {
	//TODO:: handle with exclusion data
	db.Database().Exec(fmt.Sprintf("delete from users where id = %d", userId))
}

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	route := routes.Group("/users")
	route.GET("/:userId", users.NewShowController().Show)
	route.POST("", users.NewStoreController().Store)

	return routes
}
