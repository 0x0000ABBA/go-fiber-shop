package routes

import (
	"fiber-shop/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Authorization(a *fiber.App) {
	a.Post("/authorization", controllers.Authorization)
}