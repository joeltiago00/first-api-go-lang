package users_action

import (
	"fmt"
	userRepository "github.com/joeltiago00/first-api-go-lang/infrastructure/repositories/user"
	"time"
)

//var user models.User

func Update(userId int, payload map[string]string) (bool, string) {
	if !userRepository.ExistsById(userId) {
		return true, fmt.Sprintf("User %d not found.", userId)
	}

	setUserFromPayloadUpdate(payload)
	fmt.Println(userId, user)

	userRepository.UpdateById(userId, user)

	return false, ""
}

func setUserFromPayloadUpdate(payload map[string]string) {
	user.FirstName = payload["first_name"]
	user.LastName = payload["last_name"]
	user.UpdatedAt = time.Now()
}
