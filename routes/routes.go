package routes

import (
	controller "golangrestlearn/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/v1/test", controller.Test)
}
