package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func Registration(c *fiber.Ctx) error {

	//TODO call registration service, that will return success code on success (mb email confirmation)
	//TODO (mb send jwt token on success)
	
	return errors.New("implement me")
}
