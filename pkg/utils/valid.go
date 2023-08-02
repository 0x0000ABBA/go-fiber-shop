package utils

import (
	"regexp"

	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func IsValidPassword(password string) bool {
	pattern := "^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$"

	reg := regexp.MustCompile(pattern)

	return reg.MatchString(password)
}

func IsValidName(name string) bool {

	pattern := "^[A-Za-z]+$"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(name)
}
