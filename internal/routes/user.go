package routes

import (
	"fiber-shop/internal/controllers"


	"github.com/gofiber/fiber/v2"
)

func User(a *fiber.App) {

	route := a.Group("/user")

	route.Get("/", controllers.GetAllUsers)
	route.Get("/:id", controllers.GetUserById)

}
