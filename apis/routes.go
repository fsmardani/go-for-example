package main

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/fsmardani/go-for-example/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/user", handlers.CreateUser)

	app.Get("/user/:id", handlers.ShowUser)

	app.Get("/user/:id/edit", handlers.EditFact)

	app.Delete("/user/:id", handlers.DeleteUser)
}
