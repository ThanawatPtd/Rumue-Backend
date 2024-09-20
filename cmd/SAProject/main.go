package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "SEProject",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Replace with your frontend URL
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Listen(":3001")

}
