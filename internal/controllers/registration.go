package controllers

import (
	"fiber-shop/internal/models"
	"fiber-shop/internal/services"
	"fiber-shop/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {

	u := models.User{}

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := services.SignUp(u)

	if err != nil && err == utils.ErrorUserAlreadyExists {
		return c.JSON(fiber.Map{"message": utils.ErrorUserAlreadyExists.Error()})
	}

	u, err = services.SignUp(u)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"id": u.Id})
}
