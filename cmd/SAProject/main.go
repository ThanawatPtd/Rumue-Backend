package main

import (
	"github.com/ThanawatPtd/SAProject/internal/wire"
	"github.com/ThanawatPtd/SAProject/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {

	handler := wire.InitializeHandler()

	app := fiber.New(fiber.Config{
		AppName: "SEProject",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Replace with your frontend URL
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	router.RegisterApiRouter(app, handler)

	app.Listen(":3001")

}
