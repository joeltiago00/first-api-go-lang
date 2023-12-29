package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joeltiago00/first-api-go-lang/helpers/db"
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database"
	"github.com/joeltiago00/first-api-go-lang/routes"
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

	return routes.SetupRoutes()
}
