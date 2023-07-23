package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
)

func JwtProtected(secret string) fiber.Handler {

	config := jwtware.Config{
		SigningKey:   secret,
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"message": "Corrupted or invalid jwt",
	})
}
