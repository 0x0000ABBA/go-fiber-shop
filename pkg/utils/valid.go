package utils

import (
	"regexp"
	"unicode"

	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}

func IsValidPassword(password string) bool {

	var (
		hasValidLength   = false
		hasSpecial   = false
		hasUppercase = false
		hasLowercase = false
		hasDigit     = false
		hasNoSpaces  = true
	)

	hasValidLength = len(password) > 7 && len(password) <= 32

	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			hasLowercase = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsUpper(c):
			hasUppercase = true
		case unicode.IsSpace(c):
			hasNoSpaces = false
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	return hasValidLength &&
		hasSpecial &&
		hasDigit &&
		hasUppercase &&
		hasLowercase &&
		hasNoSpaces
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func IsValidName(name string) bool {

	pattern := "^[A-Za-z]+$"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(name)
}
