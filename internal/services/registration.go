package services

import (
	"fiber-shop/internal/data/models"
	"fiber-shop/pkg/utils"
)

func SignUp(u models.User) (models.User, error) {

	// TODO validate email and password (email should be valid email and password >6 symbols, special, capital and numbers)

	if !utils.IsValidEmail(u.Email)  {
		return models.User{}, utils.ErrorInvalidEmail
	}

	if !utils.IsValidPassword(u.Password) {
		return models.User{}, utils.ErrorInvalidPassword
	}


	// TODO check if user with this email exists, if no - return error userAlreadyExists

	if UserExists(u.Email) {
		return models.User{}, utils.ErrorUserAlreadyExists
	}

	// TODO add user to db
	// TODO return new user (uuid) and 200

	

	return models.User{}, nil // TODO implement me
}
