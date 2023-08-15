package routes

import (
	"fiber-shop/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Registration(a *fiber.App) {

	a.Post("/signup", controllers.SignUp)
	
}