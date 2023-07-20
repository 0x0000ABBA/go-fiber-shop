package controllers

import (
	"fiber-shop/internal/services"
	"fiber-shop/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers returns all users
// @Description Returns all users
// @Error 500
// @Success 200
// @Router /user/ [get]
func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting users",
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUserById returns user with given ID
// @Description Returns a single user with given ID
// @Param id path string true "User ID"
// @Error 400 {string} string "Invalid user id"
// @Error 404 {string} string "User not found"
// @Error 500 {string} string "Internal server error"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	
	id := c.Params("id")

	if !utils.IsValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user id",
		})
	}

	user, err := services.GetUserById(id)

	if err != nil {
		if err == services.ErrorUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)

}