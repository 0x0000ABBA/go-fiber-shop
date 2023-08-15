package middlewares

import "github.com/gofiber/fiber/v2"

func AdminProtected() fiber.Handler {

	// GetUser and check if user is admin
	// error if not

	return func(c *fiber.Ctx) error { return nil } // TODO:implement me
}
