package routes

import (
	"fiber-shop/internal/controllers"


	"github.com/gofiber/fiber/v2"
)

func User(a *fiber.App) {

	route := a.Group("/user")

	route.Get("/", controllers.GetAllUsers) //FIXME this is kinda stupid
	route.Get("/:id", controllers.GetUserById)
	//TODO register and auth
}
