package main

import (
	"fmt"
	"github.com/joeltiago00/first-api-go-lang/config"
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database"
	"github.com/joeltiago00/first-api-go-lang/routes"
)

func main() {
	fmt.Println(config.DbConfig().User)
	database.Handle()

	routes.Handler()
}
