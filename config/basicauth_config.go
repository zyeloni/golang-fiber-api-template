package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/basicauth"
)

func LoadBasicAuthConfig() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			Cfg.ApiUser: Cfg.ApiPassword,
		},
		Realm: "Forbidden",
		Unauthorized: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}
