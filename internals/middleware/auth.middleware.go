package middleware

import (
	"strings"

	"github.com/Frank2006x/simple-bank/token"
	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is missing",
			})
		}
		fields := strings.Fields(authHeader)

		if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired access token",
			})
		}	
		c.Locals("AuthorizationPayloadKey", payload)


		return c.Next()
	}
}