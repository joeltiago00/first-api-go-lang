package requests

import (
	"fmt"
	"net/mail"
	"strconv"
	"strings"
)

var errors map[string]map[int]string

func Validate(payload map[string]string, rules map[string]string) (bool, map[string]map[int]string) {
	errors := make(map[string]map[int]string)

	for field, value := range payload {
		currentValidation := rules[field]

		for _, validations := range strings.Split(currentValidation, "|") {
			validation := strings.Split(validations, ":")

			switch validation[0] {
			case "min":
				min, _ := strconv.Atoi(validation[1])
				errors = Min(field, value, min, errors)
				break
			case "max":
				max, _ := strconv.Atoi(validation[1])
				errors = Max(field, value, max, errors)
				break
			case "email":
				errors = Email(field, value, errors)
				break
			}
		}
	}

	if len(errors) != 0 {
		return true, errors
	}

	return false, errors
}

func Min(field string, text string, quantity int, errors map[string]map[int]string) map[string]map[int]string {
	stringLength := len(text)

	if stringLength < quantity {
		errors = addError(errors, field, fmt.Sprintf("Field %s must be a least %d characters.", field, quantity))
	}

	return errors
}

func Max(field string, text string, quantity int, errors map[string]map[int]string) map[string]map[int]string {
	stringLength := len(text)

	if stringLength > quantity {
		errors = addError(errors, field, fmt.Sprintf("Field must be %d characters maximum.", quantity))
	}

	return errors
}

func Between(field string, text string, min int, max int, errors map[string]map[int]string) map[string]map[int]string {
	stringLength := len(text)

	if stringLength < min || stringLength > max {
		errors = addError(errors, field, fmt.Sprintf("Field last_name must be between %d and %d characters.", min, max))
	}

	return errors
}

func Email(field string, text string, errors map[string]map[int]string) map[string]map[int]string {
	_, err := mail.ParseAddress(text)

	if err != nil {
		errors = addError(errors, field, "Field must be a valid email.")
	}

	return errors
}

func addError(errors map[string]map[int]string, field string, error string) map[string]map[int]string {
	_, exists := errors[field]

	if exists {
		errors[field][len(errors[field])] = error

		return errors
	}

	errors[field] = make(map[int]string)
	errors[field][0] = error

	return errors
}
