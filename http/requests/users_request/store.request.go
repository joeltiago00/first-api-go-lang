package users_request

import (
	"github.com/joeltiago00/first-api-go-lang/http/requests"
)

func ValidateStore(payload map[string]string) (bool, map[string]map[int]string) {
	rules := map[string]string{
		"first_name": "min:3|max:20",
		"last_name":  "min:3|max:30",
		"email":      "email",
	}

	return requests.Validate(payload, rules)
}
