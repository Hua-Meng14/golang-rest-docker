package main

import (
	"github.com/Hua-Meng14/golang-rest-docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/fact", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}