package main

import (
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database"
	"github.com/joeltiago00/first-api-go-lang/routes"
)

func main() {
	database.Handler()

	routes.Handler()
}
