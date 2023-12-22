package users_action

import (
	userRepository "github.com/joeltiago00/first-api-go-lang/infrastructure/repositories/user"
	"github.com/joeltiago00/first-api-go-lang/models"
)

func Show(userId int) (models.User, error) {
	return userRepository.FindById(userId)
}
