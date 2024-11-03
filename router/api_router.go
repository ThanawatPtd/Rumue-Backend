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
	user.Get("/profile/id", handler.User.GetUserProfileByID)
	user.Put("/update", handler.User.UpdateUser)
	user.Put("/update/password", handler.User.UpdatePassword)
	user.Delete("/delete", handler.User.DeleteByID)

	employee := app.Group("/employee")

	employee.Post("/create", handler.Employee.CreateEmployee)
	// user.Update("/update/:id", )

	vehicle := app.Group("/vehicle")
	vehicle.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	vehicle.Post("/", handler.Vehicle.CreateVehicle)
	vehicle.Get("template", handler.Vehicle.FindTemplate)

	transaction := app.Group("/transaction")
	transaction.Use(middlewares.JwtMiddleware(config.ProvideConfig().JWTSecret))
	transaction.Put("/", handler.Transaction.UpdateTransaction)
	transaction.Get("/history", handler.Transaction.CheckHistory)    // use userID
	transaction.Get("/list", handler.Transaction.FindInsuranceToday) // transaction that pending
	transaction.Get("/:id", handler.Transaction.GetUserVehicleTransactionByID) //transactionID
	transaction.Post("/create/:id", handler.Transaction.CreateTransaction)     //vehicleID

	mile := app.Group("/mile")
	mile.Get("/:id", handler.Mile.GetMileRateByID)

	priority := app.Group("/priority")
	priority.Get("/:id", handler.Priority.GetPriorityRateByID)

	insurance := app.Group("/insurance")
	insurance.Post("", handler.Insurance.GetInsurance)
	insurance.Get("", handler.Insurance.GetInsurances)

	email := app.Group("/email")
	email.Post("/receipt", handler.Email.SendMailToAlertReceipt)
}
