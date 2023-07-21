package controllers

import "github.com/gofiber/fiber/v2"


// GetHeartbeat returns a status code to indicate server is alive
// @Description Returns a status code to indicate server is alive
// @Success 200 {string} string "Server is alive"
// @Router /heartbeat [get]
func GetHeartbeat(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Server is alive",
	})
}