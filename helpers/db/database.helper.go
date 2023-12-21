package db

import (
	"github.com/joeltiago00/first-api-go-lang/config"
	"github.com/joeltiago00/first-api-go-lang/infrastructure/database/factories"
	"gorm.io/gorm"
	"log"
)

// Database TODO:: change return to any or handle with multiple return types
func Database() *gorm.DB {
	var db *gorm.DB
	switch config.DbConfig().Connection {
	case "postgres":
		db = factories.DB
		break
	default:
		log.Panic("Invalid connection.")
	}

	return db
}
