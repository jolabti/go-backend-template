package routes

import (
	controller "golangrestlearn/controller/commons"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/v1/test/users", controller.TestUser)
}
