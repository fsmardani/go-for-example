package main

import (
	"github.com/fsmardani/go-for-example/config"
	"github.com/fsmardani/go-for-example/handlers"
	"github.com/fsmardani/go-for-example/middlewares"
	"github.com/gofiber/fiber/v2"
)


func setupRoutes(app *fiber.App) {
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	user_router := app.Group("/users")
	user_router.Post("/login", handlers.Login)

	user_router.Get("/", jwt, handlers.ListUsers)

	user_router.Post("/user", handlers.CreateUser)

	user_router.Get("/user/:id", handlers.ShowUser)

	user_router.Get("/user/:id/edit", handlers.EditUser)

	user_router.Delete("/user/:id", handlers.DeleteUser)

    image_router := app.Group("/images")

	image_router.Post("/upload", handlers.UploadFile)
}
