package routes

import (
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Swagger(a *fiber.App) {
	route := a.Group("/swagger")

	route.Get("*", swagger.HandlerDefault)
}