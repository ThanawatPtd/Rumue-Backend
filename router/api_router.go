package router

import (
	"github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/ThanawatPtd/SAProject/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouter(app *fiber.App, handler *rest.Handler) {
	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
		})
	})

	auth := app.Group("auth")

	auth.Post("/register", handler.Auth.Register)
	auth.Post("/login", handler.Auth.Login)

	user := app.Group("/user")

	user.Get("", handler.User.GetUsers)
	user.Get("/id=:id", handler.User.GetUserByID)

	user.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	user.Put("/update/id=:id", handler.User.UpdateUser)
	user.Put("/update/password/id=:id", handler.User.UpdatePassword)
	user.Delete("", handler.User.DeleteByID)
	user.Put("", handler.User.UpdateUser)

	employee := app.Group("/employee")

	employee.Post("/create", handler.Employee.CreateEmployee)
	// user.Update("/update/:id", )

	vehicle := app.Group("/vehicle")
	vehicle.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	vehicle.Post("/", handler.Vehicle.CreateVehicle)

	transaction := app.Group("/transaction")
	transaction.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	transaction.Post("/create/id=:id", handler.Transection.CreateTransaction)
}
