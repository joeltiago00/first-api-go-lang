package factories

import (
	"github.com/joeltiago00/first-api-go-lang/helpers/db"
	"github.com/joeltiago00/first-api-go-lang/models"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (r *UserFactory) Create() models.User {
	user := models.User{FirstName: "Testing", LastName: "Testing", Email: "test@mail.com"}

	db.Database().Create(&user)

	return user
}
