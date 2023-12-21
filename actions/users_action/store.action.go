package users_action

import (
	userRepository "github.com/joeltiago00/first-api-go-lang/infrastructure/repositories/user"
	"github.com/joeltiago00/first-api-go-lang/models"
)

var user models.User

func Store(payload map[string]string) models.User {
	setUserFromPayload(payload)

	return userRepository.Store(user)
}

func setUserFromPayload(payload map[string]string) {
	user.FirstName = payload["first_name"]
	user.LastName = payload["last_name"]
	user.Password = payload["password"]
	user.Email = payload["email"]
}
