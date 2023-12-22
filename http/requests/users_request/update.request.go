package users_request

import (
	"github.com/joeltiago00/first-api-go-lang/http/requests"
)

func ValidateUpdate(payload map[string]string) (bool, map[string]map[int]string) {
	return requests.Validate(payload, map[string]string{
		"first_name": "min:3|max:20",
		"last_name":  "min:3|max:30",
	})
}
