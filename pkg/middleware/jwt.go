package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
)

func JwtProtected(secret string) func(*fiber.Ctx) error {

	config := jwtware.Config{
		SigningKey:   secret,
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	return nil
}
