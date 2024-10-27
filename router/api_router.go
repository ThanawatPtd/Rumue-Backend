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

	user.Get("", handler.User.GetUsers) //who use this na

	user.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	user.Get("/id", handler.User.GetUserByID)
	user.Put("/update", handler.User.UpdateUser)
	user.Put("/update/password", handler.User.UpdatePassword)
	user.Delete("/delete", handler.User.DeleteByID)

	employee := app.Group("/employee")

	employee.Post("/create", handler.Employee.CreateEmployee)
	// user.Update("/update/:id", )

	vehicle := app.Group("/vehicle")
	vehicle.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	vehicle.Post("/", handler.Vehicle.CreateVehicle)

	transaction := app.Group("/transaction")
	transaction.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))

	// Define routes
	transaction.Put("/", handler.Transection.UpdateTransaction)
	transaction.Get("/history", handler.Transection.CheckHistory)    // use userID
	transaction.Get("/list", handler.Transection.FindInsuranceToday) // transaction that pending
	transaction.Get("/:id", handler.Transection.GetUserVehicleTransactionByID)
	transaction.Post("/create/:id", handler.Transection.CreateTransaction)
}
