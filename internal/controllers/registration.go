package controllers

import (
	"errors"
	"fiber-shop/internal/data/models"
	"fiber-shop/internal/services"
	"fiber-shop/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func Registration(c *fiber.Ctx) error {

	u := models.User{}

	if err := c.BodyParser(u); err != nil {
		return err
	}

	_, err := services.SignUp(u)

	if err != nil && err == utils.ErrorUserAlreadyExists {
		return c.JSON(fiber.Map{"message": utils.ErrorUserAlreadyExists.Error()})
	}


	//TODO call registration service, that will return success code on success (mb email confirmation)
	//TODO (mb send jwt token on success)
	
	return errors.New("implement me")
}
