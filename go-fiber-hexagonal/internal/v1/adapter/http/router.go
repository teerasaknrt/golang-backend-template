package http

import (
	"go-fiber-hexagonal/internal/v1/adapter/http/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, uh *handler.UserHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/users/:id", uh.GetUser)
}
