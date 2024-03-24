package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Hua-Meng14/golang-rest-docker/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}