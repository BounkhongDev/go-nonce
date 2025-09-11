package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-nonce/config"
	"go-nonce/database"
	"go-nonce/logs"
	"go-nonce/routes"
	"go-nonce/src/controllers"
	"go-nonce/src/services"
	"log"
)

func main() {

	//connect database
	postgresConnection, err := database.PostgresConnection()
	if err != nil {
		logs.Error(err)
		return
	}

	//connect redis
	database.InitRedis()

	//basic structure
	// connect route
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(cors.New())

	//example routes
	exampleService := services.NewExampleService(postgresConnection)
	exampleController := controllers.NewExampleController(exampleService)
	newRoute := routes.NewFiberRoutes(
		//new web controller
		exampleController,
	)
	// newRoute.Install(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env("app.port"))))
}
