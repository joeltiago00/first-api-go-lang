package tests

import (
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database"
	"testing"
)

// FeatureSetup Make application setup to run tests
func FeatureSetup(test *testing.T) {
	// Set env vars for tests
	test.Setenv("APP_ENV", "testing")

	database.Handler()
}
