package handlers

import "github.com/gofiber/fiber/v3"

func Homepage(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
