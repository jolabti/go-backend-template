package controller

import (
	"context"
	model "golangrestlearn/model"
	"log"

	helper "golangrestlearn/helpers"

	configurationDatabase "golangrestlearn/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// func Test(c *fiber.Ctx) error {

// 	resMapping := []model.User{
// 		{
// 			ID:       1,
// 			UserName: "user1",
// 			Email:    "asdsd",
// 		},
// 		{
// 			ID:       2,
// 			UserName: "user2",
// 			Email:    "asdsd2",
// 		},
// 	}
// 	res := helper.ResultMessage{
// 		Code:    200,
// 		Data:    resMapping,
// 		Message: "",
// 	}

// 	return c.JSON(res)
// }

func TestUser(c *fiber.Ctx) error {

	var ctx = context.Background()

	db, err := configurationDatabase.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer csr.Close(ctx)

	result := make([]model.User, 0)

	for csr.Next(ctx) {
		var row model.User
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	return helper.ResponseMessage(c, 200, "Get Data Succesfully", result)
}
