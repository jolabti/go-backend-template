package main

import (
	"fmt"
	"log"
	"os"

	routes "golangrestlearn/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// creating fiber instance
	app := fiber.New()

	godotenv.Load()

	// listing the host setup for creating connection to database

	createConnection()

	//create http handler
	// app.Use(app)
	routes.Setup(app)

	app.Listen(":3001")
}

func createConnection() {

	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_DBUSER")
	dbname := os.Getenv("POSTGRESQL_DBNAME")
	password := os.Getenv("POSTGRESQL_DBPASSWORD")

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	var db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error database connection")
		log.Println("This is an info message." + db.Name())
	}

	DB = db

	// fmt.Printf("Connection to database")
}
