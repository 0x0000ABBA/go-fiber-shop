package routes

import (
	"fiber-shop/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Product(a *fiber.App) {
	a.Group("/product")

	a.Get("/:id", controllers.GetProductById)
	a.Get("/all", controllers.GetAllProducts)
	a.Get("/page/:page", controllers.GetProductPage)
}