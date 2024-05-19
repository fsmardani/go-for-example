package main

import (
	"github.com/fsmardani/go-for-example/handlers"
	"github.com/fsmardani/go-for-example/middlewares"
	"github.com/fsmardani/go-for-example/config"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	app.Post("/login", handlers.Login)
	
	app.Get("/", handlers.ListUsers)

	app.Post("/user", jwt, handlers.CreateUser)

	app.Get("/user/:id", handlers.ShowUser)

	app.Get("/user/:id/edit", handlers.EditUser)

	app.Delete("/user/:id", handlers.DeleteUser)
}
