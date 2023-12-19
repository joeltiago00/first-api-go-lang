package database

import (
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database/factories"
	"log"
)

func DatabaseFactoryHandle(databaseConnection string) {
	switch databaseConnection {
	case "postgres":
		factories.Postgres()
	default:
		log.Panic("Database connection " + databaseConnection + " are not implemented yet.")
	}
}
