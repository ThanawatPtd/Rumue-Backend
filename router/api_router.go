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

	user := app.Group("/user")

	user.Get("", handler.User.GetUsers)
	user.Get("/id=:id", handler.User.GetUserByID)

	user.Post("/register", handler.User.Register)
	user.Post("/login", handler.User.Login)
	user.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	user.Delete("", handler.User.DeleteByID)
	user.Put("", handler.User.UpdateUser)

	employee := app.Group("/employee")

	employee.Post("/create", handler.Employee.CreateEmployee)
	// user.Update("/update/:id", )

	vehicle := app.Group("/vehicle")

	vehicle.Post("/", handler.Vehicle.CreateVehicle)
}
