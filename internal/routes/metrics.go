package routes

import (
	"fiber-shop/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Metrics(a *fiber.App) {
	route := a.Group("/metrics")

	route.Get("/heartbeat", controllers.GetHeartbeat)
}
