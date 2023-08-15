package services

import (
	"fiber-shop/internal/models"
	"fiber-shop/pkg/utils"
)

func SignUp(u models.User) (models.User, error) {

	if !utils.IsValidEmail(u.Email) {
		return models.User{}, utils.ErrorInvalidEmail
	}

	if !utils.IsValidPassword(u.Password) {
		return models.User{}, utils.ErrorInvalidPassword
	}

	if UserExists(u.Email) {
		return models.User{}, utils.ErrorUserAlreadyExists
	}

	// TODO:add user to db
	// TODO:return new user (uuid) and 200

	return models.User{}, nil // TODO:implement me
}
