package controller

import (
	middleware "golangrestlearn/middleware"
	model "golangrestlearn/model"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {

	resMapping := []model.User{
		{
			ID:       1,
			UserName: "user1",
			Password: "123213",
			Email:    "asdsd",
		},
		{
			ID:       2,
			UserName: "user2",
			Password: "123212",
			Email:    "asdsd2",
		},
	}
	res := middleware.ResponseWrapper{
		Data:       resMapping,
		StatusCode: 200,
	}

	return c.JSON(res)
}
