package commons

import (
	"context"
	model "golangrestlearn/model"
	"log"

	helper "golangrestlearn/helpers"

	configurationDatabase "golangrestlearn/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUser(c *fiber.Ctx) error {

	var ctx = context.Background()

	// redisConnector := configurationDatabase.NewRedisWrapper(0)

	// redisConnector.Set("testcoffee", "hot", 60*time.Second)

	db, err := configurationDatabase.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("content-section").Find(ctx, bson.M{})
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
