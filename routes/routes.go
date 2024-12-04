package routes

import (
	"github.com/gofiber/fiber/v3"
	"golang-fiber-api-template/handlers"
)

func LoadRoutes(router fiber.Router) {
	router.Get("/home", handlers.Homepage)
}
