package main

import (
	"github.com/gofiber/fiber/v3"
	"golang-fiber-api-template/config"
	"golang-fiber-api-template/internal/database"
	"golang-fiber-api-template/internal/routes"
	"log"
)

func main() {
	config.LoadConfig()
	err := database.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}

	database.RunMigrations()

	app := fiber.New()
	app.Use(config.LoadBasicAuthConfig())

	router := app.Group("/api")
	routes.LoadRoutes(router)

	err = app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
		return
	}
}
