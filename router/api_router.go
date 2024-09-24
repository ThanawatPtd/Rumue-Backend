package router

import (

	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
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


	user.Get("",handler.User.GetUsers)
	user.Get("/id=:id",handler.User.GetUserByID)
	user.Get("/email=:email",handler.User.GetUserByEmail)

	user.Post("/create", handler.User.CreateUser)

	user.Delete("/delete/id=:id", handler.User.DeleteByID)


	employee := app.Group("/employee")

	employee.Post("/create",handler.Employee.CreateEmployee)
	// user.Update("/update/:id", )

}
