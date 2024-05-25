package main

import (
	"github.com/fsmardani/go-for-example/config"
	"github.com/fsmardani/go-for-example/handlers"
	"github.com/fsmardani/go-for-example/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	app.Post("/login", handlers.Login)

	app.Get("/", jwt, handlers.ListUsers)

	app.Post("/user", handlers.CreateUser)

	app.Get("/user/:id", handlers.ShowUser)

	app.Get("/user/:id/edit", handlers.EditUser)

	app.Delete("/user/:id", handlers.DeleteUser)
}
