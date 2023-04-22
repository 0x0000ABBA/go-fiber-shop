package controllers

import "github.com/gofiber/fiber/v2"

// Heartbeat is a simple heartbeat controller
// @Description Simple heartbeat controller
// @Success 200
// @Router /metrics/heartbeat [get]
func GetHeartbeat(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNoContent).SendString("")
}