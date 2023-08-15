package routes

import (
	"fiber-shop/internal/controllers"
	"fiber-shop/internal/middlewares"
	"fiber-shop/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func Product(a *fiber.App) {
	a.Group("/product")

	a.Get("/:id", controllers.GetProductById)
	a.Get("/all", controllers.GetAllProducts)
	a.Get("/page/:page", controllers.GetProductPage)

	a.Post("", middleware.JwtProtected("JWT_SECRET"), middlewares.AdminProtected(), controllers.PostProduct) // FIXME:middleware should not depend on secret
	// mb move config.GetConfigValue inside jwtProtected
}
