package factories

import (
	"fmt"
	"github.com/joeltiago00/first-api-go-lang/config"
	"github.com/joeltiago00/first-api-go-lang/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Postgres() {
	dbConfig := config.DbConfig()

	DB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.Port,
	)))

	if err != nil {
		log.Panic("Postgres connection error.")
	}

	DB.AutoMigrate(&models.User{})
}
