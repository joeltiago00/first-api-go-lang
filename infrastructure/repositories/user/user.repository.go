package user

import (
	"github.com/joeltiago00/first-api-go-lang/helpers/db"
	"github.com/joeltiago00/first-api-go-lang/models"
)

func ExistsByEmail(email string) bool {
	var user models.User

	db.Database().Where(models.User{Email: email}).First(&user)

	if user.ID == 0 {
		return false
	}

	return true
}

func ExistsById(userId int) bool {
	var user models.User

	db.Database().Where(models.User{ID: userId}).First(&user)

	if user.ID == 0 {
		return false
	}

	return true
}

func Store(user models.User) models.User {
	db.Database().Create(&user)

	return user
}

func UpdateById(userId int, user models.User) {
	db.Database().Where(&models.User{ID: userId}).UpdateColumns(user)
}
